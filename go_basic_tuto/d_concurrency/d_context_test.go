package d_concurrency

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
*/
