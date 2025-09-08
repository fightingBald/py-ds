package d_concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/*
1. select { ... }

Go 的并发控制器，可以同时监听多个 channel 操作。
这里有两个 case：
等待 ch 的响应。
等待 time.After 的超时信号。
谁先就绪，谁就执行。*/
// * select：多路复用
func TestSelect(t *testing.T) {
	ch := make(chan string)

	// 模拟外部调用晚点儿返回
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "OK"
	}()

	select {
	case resp := <-ch:
		fmt.Println("resp:", resp)
	case <-time.After(100 * time.Millisecond): // 100ms 后发个信号过来
		fmt.Println("timeout, fallback") // 100ms 超时先触发
	}
}

/*
*WaitGroup 就是一个线程安全的计数器。
Add(n) 就是 +n
Done() 就是 -1
Wait() 就是 阻塞直到计数器归零

老板（主协程）：wg.Add(工人数)
每个工人：干完 -> wg.Done()
老板：wg.Wait()，等所有人报到 -> 继续走
*/

// 全局 WaitGroup，母协程和子协程都用它
var wgTest sync.WaitGroup

// 子协程逻辑
func worker(id int) {
	defer wgTest.Done() // -1
	fmt.Printf("Worker %d is working\n", id)
}

func TestWaitGroup(t *testing.T) {
	for i := 1; i <= 3; i++ {
		wgTest.Add(1) // +1
		go worker(i)  // 开子协程去干活
	}

	wgTest.Wait()
	fmt.Println("All workers finished, boss goes home")
}
