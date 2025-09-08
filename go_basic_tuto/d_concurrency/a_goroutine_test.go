package d_concurrency

import (
	"testing"
)

func TestGoroutine(t *testing.T) {
	ch := make(chan string)
	go func() { ch <- "ping" }() // 没人收就会卡在这
	//你的代码里的确阻塞了，只是测试函数先结束了，你没看见而已。go test 不会“替你等”那个匿名 goroutine；
	//测试一返回，进程就收工，剩下那个在 ch <- "ping" 上卡住的 goroutine 直接被进程退出一起带走。
	//msg := <-ch
	//fmt.Println(msg)
	select {} // 主 goroutine 永久阻塞
	// fatal error: all goroutines are asleep - deadlock!

}
