package d_concurrency

import (
	"context"
	"fmt"
	"sync"
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

// 工具：安全地从 ctx 等待，返回是否被取消
func waitUntilCanceled(ctx context.Context, max time.Duration) bool {
	timer := time.NewTimer(max)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return true
	case <-timer.C:
		return false
	}
}

// 1) 取消传播：父协程取消，子协程立刻收摊
/*

父 ctx cancel 后，所有拿着这个 ctx 的 goroutine 都必须立刻感知并退出。
教你把退出条件写进 select { case <-ctx.Done(): ... }，而不是写在循环外面“碰运气”。

你要学到：
“广播式取消”是官方姿势，不要自造一堆 done channel。
用原子状态/断言确保 worker 真退了，别留内存泄漏。


*/
var stopped atomicBool

func workerTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "stopped:", ctx.Err())
			return
		default:
			fmt.Println(name, "working...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func TestCancelPropagation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// 父工人
	go func() {
		// 子工人
		go workerTask(ctx, "child-1")
		go workerTask(ctx, "child-2")
	}()

	time.Sleep(1 * time.Second)
	cancel() // 广播信号：全家收工
	time.Sleep(500 * time.Millisecond)
}

// 2) 超时：到点不等了，返回 context.DeadlineExceeded
func TestTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(200 * time.Millisecond): // 假装下游卡住很久
		t.Fatal("should have timed out")
	case <-ctx.Done():
		if ctx.Err() != context.DeadlineExceeded {
			t.Fatalf("expected DeadlineExceeded, got %v", ctx.Err())
		}
	}
}

// 3) 截止时间：指定绝对时间点
func TestDeadline(t *testing.T) {
	deadline := time.Now().Add(60 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	<-ctx.Done()
	if ctx.Err() != context.DeadlineExceeded {
		t.Fatalf("expected DeadlineExceeded, got %v", ctx.Err())
	}
}

// 4) 传值（轻量元数据）：trace id 往下传
type ctxKey string

func TestWithValue(t *testing.T) {
	const traceKey ctxKey = "traceID"
	traceID := "req-12345"

	ctx := context.WithValue(context.Background(), traceKey, traceID)

	got, _ := ctx.Value(traceKey).(string)
	if got != traceID {
		t.Fatalf("expected %s, got %s", traceID, got)
	}
	// 反例：取不存在的 key 返回 nil
	if v := ctx.Value(ctxKey("nope")); v != nil {
		t.Fatal("unexpected value for missing key")
	}
}

// 5) 配合 WaitGroup：先 cancel 再等，干净退出
func TestCancelWithWaitGroup(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	startWorkers := func(n int) {
		wg.Add(n)
		for i := 0; i < n; i++ {
			go func(id int) {
				defer wg.Done()
				for {
					select {
					case <-ctx.Done():
						// 收尾动作放这里，比如释放连接、flush 缓冲等
						return
					default:
						time.Sleep(5 * time.Millisecond)
					}
				}
			}(i)
		}
	}

	startWorkers(5)
	time.Sleep(20 * time.Millisecond) // 让大家跑两步
	cancel()                          // 广播取消
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()

	select {
	case <-done:
		// ok
	case <-time.After(300 * time.Millisecond):
		t.Fatal("workers did not exit after cancel")
	}
}

// 6) 给外部调用包一层优雅超时（最贴近真实项目）
func doSlowOp(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func TestWrapOperationWithTimeout(t *testing.T) {
	// 业务层给下游设 80ms 超时
	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer cancel()

	err := doSlowOp(ctx, 200*time.Millisecond)
	if err == nil || err != context.DeadlineExceeded {
		t.Fatalf("expected DeadlineExceeded, got %v", err)
	}

	// 快速路径应该不报错
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	if err := doSlowOp(ctx2, 10*time.Millisecond); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
}

/******** 小工具：原子 bool，避免引第三方库 ********/
type atomicBool struct {
	mu sync.Mutex
	v  bool
}

func (a *atomicBool) set(b bool) { a.mu.Lock(); a.v = b; a.mu.Unlock() }
func (a *atomicBool) get() bool  { a.mu.Lock(); defer a.mu.Unlock(); return a.v }
