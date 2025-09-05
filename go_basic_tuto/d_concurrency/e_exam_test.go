package d_concurrency

import (
	"sync"
	"testing"
)

// 我被多次要求编写一个程序，使用两个 Goroutine 依次打印从 1 到 10 的数字
// 为了使 goroutine 按顺序运行，需要阻止执行或等待来自另一个 goroutine 的信号来恢复。
// 假设您必须创建一个程序，使用 2 个 goroutine 按顺序打印从 1 到 5 的数字，其中 1 个 goroutine 打印奇数，另一个打印偶数。
// 我们将使用通道来通知 Goroutine 彼此的更新。对于两个 Goroutine，创建两个布尔类型的通道：偶数通道和奇数通道。
//
// 奇数号的 Goroutine 会等待 even 通道，一旦收到信号，它就会恢复并打印奇数号，然后通知使用 odd 通道的偶数号 Goroutine，
// 并等待偶数号 Goroutine 发出信号。偶数号的 Goroutine 一旦从 odd 通道收到信号，就会恢复并打印偶数号，
// 然后通知奇数号 Goroutine。如此循环，直到它们到达循环末尾。
var wg sync.WaitGroup

func TestExam(t *testing.T) {
	wg.Add(2)
	oddChan := make(chan bool)
	evenChan := make(chan bool)

	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i += 2 {
			<-oddChan
			println(i)
			evenChan <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 5; i += 2 {
			<-evenChan
			println(i)
			oddChan <- true
		}
	}()

	//启动奇数号 Goroutine
	oddChan <- true
	wg.Wait()
}
