package d_concurrency

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
是的，context 本质上就是：
一个 goroutine 往下游所有 goroutine 传递“活着/死了、还能干多久、小标签”的机制。

context 就是一个信号广播器，能让一个 goroutine 给它的所有子 goroutine（以及孙子、重孙子……只要链路上传递了同一个 ctx）发消息：
“兄弟们，该停工了” → ctx.Done()
“你们最多干到几点” → ctx.Deadline()
“这是我们这趟活的工单号” → ctx.Value(key)

为什么要这么设计 (避免孤儿 goroutine)
在 Go 里，你随便能开一堆 goroutine，但如果没有统一的控制，主 goroutine 结束了，子 goroutine 可能还挂在那儿等 IO，结果内存泄漏、死锁、乱七八糟。
context 解决了这几个问题：
* 广播取消：一个 cancel 信号能层层传递，不需要你自己维护一堆 channel。
* 超时 / 截止时间：保证 goroutine 不会无限期挂着。
* 共享请求元数据：traceID、userID 之类的小信息沿调用链传到底。


context 是单向的 ——它就是从 父 goroutine 往子孙 goroutine 广播“存活/超时/取消”这些状态的。
由下往上呢？几乎没戏。
1. 为什么不能由下往上传
context 的设计目标：取消、超时、request-scoped metadata。


不是，一个后端 API 不会把一堆请求塞到同一个 context 里乱炖。每个请求都有自己的 context，从入口一生成就跟着这条请求一路往下传，互不串味。你担心的那种“所有人共用一个 context，auth 全糊一起”的场景，只会出现在屎山项目里。
正确认知
每个请求一个 context：HTTP 框架会把它挂在 *http.Request 上，r.Context() 拿到的就是这次请求的那份。
context 是不可变的“链”：你不会“修改原 context”，而是基于它派生新的（WithCancel/WithTimeout/WithValue），像在链上加一节。父子之间取消会级联，但数据不混。
Auth/Trace 这类 request-scoped 信息放在该请求自己的 context 里很正常；服务级资源（DB 连接池、配置）走依赖注入或 struct 字段，别往 ctx 里倒垃圾。


几条硬规范，别作死

函数签名第一个参数放 ctx：func Do(ctx context.Context, ...)，好传好用。
只放 request-scoped 小数据：traceID/userID/locale 这类。不要塞 DB、配置、大对象。
派生要记得 cancel()：WithTimeout/WithCancel 用了就 defer cancel()，不然计时器泄露。
key 要用自定义类型：避免字符串撞车：type ctxKey string，并且 key 常量不导出。
结论：一个 API 处理成千上万请求没问题，每个请求都有自己的 context，装自己的认证元数据和倒计时，
谁都不跟谁串号。你要是把整站共用一个 context，那不叫工程，那叫事故现场。


常见用法清单
一键取消一串 goroutine：WithCancel 往下传，子任务 select <-ctx.Done() 收到就停。
给耗时操作套超时：WithTimeout 包住 DB/HTTP/RPC 调用，超时自动打断。
传递一点点元数据：WithValue 放 requestID、userID、locale 这类轻量信息。
服务优雅下线：收到信号时 cancel()，让所有 handler/worker 有序退出。
配合 WaitGroup：先取消再 wg.Wait()，保证不留孤儿协程。


1. 为啥要有 context

想象场景：

你起了一堆 goroutine 下载、算账、查库。

用户点了“取消”，或者 HTTP 请求超时了。

你得通知所有 goroutine：别干了，赶紧收摊。

如果没 context，你只能自己造轮子：搞一堆 done channel，传来传去，乱成一坨。
Go 官方直接给你塞了个统一的东西：context.Context。

2. context 核心功能

取消（cancelation）
	一处取消，所有 goroutine 都能感知。
	避免 goroutine 泄漏（主任务没了，子任务还在偷偷干活）。

超时/截止时间（timeout / deadline）
	到点了自动取消，省心。

传值（key-value）
	可以带点元数据（比如 trace ID、用户 ID）。
	但注意：不是给你当 map 用的，官方说只能放“跨 API 边界的请求级参数”。
*/

/*************** 共用的 worker 和小工具 ***************/
// 周期性干活，直到 ctx 结束。结束时向 stopped channel 发信号。
/*
周期性干活（用 time.Ticker 每隔一段时间触发）。
一旦 <-ctx.Done() 可读，立刻收摊退出。
退出前，往 stopped 里发个空 struct，通知测试：“我确实死了”。
*/
func workerLoop(ctx context.Context, name string, tick time.Duration, stopped chan<- struct{}) {
	defer func() { stopped <- struct{}{} }() //退出前，往 stopped 里发个空 struct，通知测试：“我确实死了”。
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "stopped:", ctx.Err())
			return
		case <-ticker.C:
			// 模拟一轮活
			_ = name // 这里不打印，避免噪音
		}
	}
}

// 读取 ctx 里的值，传回结果（演示 WithValue 传递）

/*
mustWithin
测试工具函数：在给定时间 d 内，反复检查某个条件函数 fn() 是否变 true。
如果超时还没变 true，就 t.Fatal(msg) 直接失败。
相当于测试里的“轮询断言
*/
func mustWithin(t *testing.T, fn func() bool, d time.Duration, msg string) {
	dead := time.After(d)
	for {
		select {
		case <-dead:
			t.Fatal(msg)
		default:
			if fn() {
				return
			}
			time.Sleep(2 * time.Millisecond)
		}
	}
}

/******************** 1) 取消传播 ********************/

func TestCancelPropagation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	stop1 := make(chan struct{}, 1)
	stop2 := make(chan struct{}, 1)

	go workerLoop(ctx, "child-1", 20*time.Millisecond, stop1)
	go workerLoop(ctx, "child-2", 20*time.Millisecond, stop2)

	time.Sleep(80 * time.Millisecond) // 让他们真在跑
	cancel()                          // 广播“收工”

	mustWithin(t,
		func() bool {
			select {
			case <-ctx.Done():
				return true
			default:
				return false
			}
		},
		150*time.Millisecond,
		"context not canceled in time",
	)

	mustWithin(t, func() bool {
		select {
		case <-stop1:
			return true
		default:
			return false
		}
	}, 150*time.Millisecond, "child-1 did not stop")

	mustWithin(t, func() bool {
		select {
		case <-stop2:
			return true
		default:
			return false
		}
	}, 150*time.Millisecond, "child-2 did not stop")
}

/*
解说与观察：
- 你会看到：两个 worker 在 cancel() 后很快退出，ctx.Done() 变为可读。
- 结论：父 cancel 会“广播式”传到所有拿着同一个 ctx 的 goroutine。
- 要点：在循环里 select `<-ctx.Done()` 才能及时响应，不要放循环外面碰运气。
*/

/******************** 2) 超时控制（WithTimeout） ********************/
//教你用 context.WithTimeout 给下游“IO风格任务”设硬超时，并验证两件事：
//超时后返回 context.DeadlineExceeded，2) 返回时间大致在你设的 80ms 附近，而不是等到 300ms 才磨蹭完。
// 一个“IO”任务：要么等到 workDur 完成，要么被 ctx 打断。
func workerIO(ctx context.Context, workDur time.Duration) error {
	select {
	case <-ctx.Done():
		fmt.Printf("Exceed 80*time.Millisecond, return [%v] \n", ctx.Err()) //context deadline exceeded
		return ctx.Err()
	case <-time.After(workDur):
		return nil
	}
}

func TestWithTimeout(t *testing.T) {
	//WithTimeout(80ms) 给调用包了个 SLA：80ms 内不完成就算你输。
	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond)

	defer cancel()

	start := time.Now()

	err := workerIO(ctx, 300*time.Millisecond)

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("want DeadlineExceeded, got %v", err)
	}
	elapsed := time.Since(start)
	if elapsed < 70*time.Millisecond || elapsed > 200*time.Millisecond {
		t.Fatalf("timeout fired oddly, elapsed=%v", elapsed)
	}
}

/*
解说与观察：
- 你会看到：任务没做完就被打断，并返回 DeadlineExceeded。
- 结论：WithTimeout 本质就是在 ctx 上挂了 time.Timer，到点自动 cancel。
- 要点：即便超时了，也要 defer cancel()，不然 timer 和子节点可能泄漏。
*/

/******************** 3) 截止时间（WithDeadline） ********************/
//- 你会看到：在指定时刻自动结束，和 WithTimeout 作用等价，接口不同。
//- 结论：适合统一以“绝对时间点”管理多段工作（比如全链路硬截止）。
func TestWithDeadline(t *testing.T) {
	deadline := time.Now().Add(60 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()

	err := workerIO(ctx, 300*time.Millisecond)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("want DeadlineExceeded, got %v", err)
	}
}

/*
解说与观察：
- 你会看到：在指定时刻自动结束，和 WithTimeout 作用等价，接口不同。
- 结论：适合统一以“绝对时间点”管理多段工作（比如全链路硬截止）。
*/

/******************** 4) 请求级元数据（WithValue） ********************/
type ctxKey struct{}

func workerReadValue(ctx context.Context, out chan<- string) {
	v, _ := ctx.Value(ctxKey{}).(string)
	out <- v //把拿到的 traceID 发到通道里。
}

// parent := context.WithValue(context.Background(), ctxKey{}, "trace-abc123")
// WithValue 会生成一个 新的 context，内部多存了一条键值对：
// key = ctxKey{}，value = "trace-abc123"。
// 官方强烈建议：
// 用你自己定义的私有类型做 key，避免和别人打架
// 这就是 type ctxKey struct{} 的作用：一个不会和别人重复的独立类型。
// parent 里就带上了这个 trace id。
// 为啥“再包一层”？
// 因为 context 在 Go 里是不可变的，每次加数据或加超时，都会返回一个新对象，内部指向父 context。这样形成一个链表式的“上下文链”。
func TestWithValuePropagation(t *testing.T) {
	parent := context.WithValue(context.Background(), ctxKey{}, "trace-abc123")
	child, cancel := context.WithCancel(parent)
	defer cancel()

	out := make(chan string, 1)
	go workerReadValue(child, out)

	select {
	case got := <-out: //从通道里读出 traceID
		if got != "trace-abc123" {
			t.Fatalf("value lost, got=%q", got)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("reader blocked")
	}
}

/*
解说与观察：
- 你会看到：子协程能读到父 ctx 写入的 traceID。
- 结论：WithValue 通过父子链向上查值，适合传“轻量”请求元数据。
- 要点：key 用自定义私有类型，避免冲突；别往里面塞 DB 连接这类重物。
*/

/******************** 5) ctx 结束时的清理（AfterFunc） ********************/

func TestAfterFunc(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{}, 1)

	stop := context.AfterFunc(ctx, func() {
		// 清理动作，例如：释放缓存、打点上报
		done <- struct{}{}
	})
	defer stop()

	cancel() // 触发 after-func

	select {
	case <-done:
		// ok
	case <-time.After(150 * time.Millisecond):
		t.Fatal("after-func did not run")
	}
}

/*
解说与观察：
- 你会看到：cancel() 后，注册的回调被执行。
- 结论：AfterFunc 允许把收尾逻辑挂在 ctx 生命周期上。
- 要点：函数返回的 stop() 可以尝试阻止未执行的回调，记得在合适时机调用。
*/

/******************** 6) 取消原因（WithCancelCause / Cause） ********************/

func TestWithCancelCause(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	cause := errors.New("auth failed")
	cancel(cause)

	<-ctx.Done()
	if !errors.Is(context.Cause(ctx), cause) {
		t.Fatalf("want cause=%v, got %v", cause, context.Cause(ctx))
	}
}

/*
解说与观察：
- 你会看到：不仅知道“被取消”，还能拿到具体原因。
- 结论：比只有 Canceled/DeadlineExceeded 更可观测，排错更省心。
*/

/******************** 7) 切断取消链（WithoutCancel） ********************/

func TestWithoutCancel(t *testing.T) {
	parent := context.WithValue(context.Background(), ctxKey{}, "req-999")
	parent, parentCancel := context.WithTimeout(parent, 60*time.Millisecond)
	defer parentCancel()

	child := context.WithoutCancel(parent) // 不继承取消，只继承 Value

	// value 应保留
	if got := child.Value(ctxKey{}).(string); got != "req-999" {
		t.Fatalf("value not preserved, got=%v", got)
	}

	// 等父亲超时
	<-parent.Done()

	// child 不应该被取消
	select {
	case <-child.Done():
		t.Fatal("child should not be canceled by parent")
	case <-time.After(40 * time.Millisecond):
		// ok
	}
}

/*
解说与观察：
- 你会看到：父亲死了，child 仍活着，但它仍然携带 Value。
- 结论：当你不想被上游 cancel 误伤、但还想带着元数据下沉时，用它。
- 要点：别滥用，切断取消链会破坏上游的“全链路生杀权”。
*/

/******************** 8) HTTP 中的 ctx（请求超时） ********************/

func TestHTTPWithContextTimeout(t *testing.T) {
	// 模拟慢服务
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		fmt.Fprintln(w, "ok")
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{} // 不用 client.Timeout，纯靠 ctx

	_, err = client.Do(req)
	if err == nil {
		t.Fatal("expected cancellation error, got nil")
	}
}

/*
解说与观察：
- 你会看到：客户端 80ms 超时，HTTP 请求中断，Do 返回错误。
- 结论：HTTP、数据库等标准库/驱动普遍支持 Context，拿着它就能“随时掐电”。
- 要点：server 端 Handler 里也能读 r.Context()，向下游传递取消信号。
}

/******************** 9) 循环里的优雅退出（worker 版） ********************/

func TestWorkerLoopWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stopped := make(chan struct{}, 1)
	go workerLoop(ctx, "loop", 30*time.Millisecond, stopped)

	time.Sleep(100 * time.Millisecond)
	cancel()

	select {
	case <-stopped:
		// ok
	case <-time.After(150 * time.Millisecond):
		t.Fatal("worker did not stop in time")
	}
}

/*
解说与观察：
- 你会看到：循环型 worker 能在 cancel 后很快结束。
- 结论：select + ticker 是循环任务的惯用骨架，别写死循环不看 Done。
- 要点：ticker.Stop() 别忘了，不然测试多了你自己卡死。
*/
