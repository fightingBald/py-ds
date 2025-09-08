package d_concurrency

import (
	"fmt"
	"testing"
	"time"
)

//所有的 Goroutine 之间都是相互隔离的，彼此之间并不知道对方的存在。Goroutine 使用 channel 作为桥梁来相互通信。
//速记：channel 基本操作
//创建：ch := make(chan T) / 缓冲：make(chan T, n)
//发：ch <- v（可能阻塞）
//收：x := <-ch（可能阻塞）
//关：close(ch)（只能发送方关）
//遍历：for v := range ch { ... }（直到关闭

// * 1) 无缓冲 channel：天然同步点（“你不收我就卡着”）
func TestChannel(t *testing.T) {

	ch := make(chan string)

	// 生产者：晚点儿发
	go func() {
		time.Sleep(10000 * time.Millisecond)
		ch <- "ping" // 没人收就一直阻塞在这
	}()

	msg := <-ch // 主 goroutine 在这等，直到对方把东西塞进来

	fmt.Println(msg) // ping
}

//* 2) 有缓冲 channel：

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 2) // 容量=2

	ch <- 1 // 不阻塞
	ch <- 2 // 不阻塞
	// ch <- 3 // 现在会阻塞，除非开始有人收

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
}

//* 3) 生产者关闭 + 消费者 range：优雅收尾

func TestChannelClose(t *testing.T) {
	ch := make(chan int)

	// 生产者
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch) // 只有发送方关；关了表示“以后没货了”
	}()

	// 消费者
	for v := range ch { // 读到通道关闭自动退出
		fmt.Println("got:", v)
	}
}
