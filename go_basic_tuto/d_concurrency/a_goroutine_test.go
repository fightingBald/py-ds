package d_concurrency

import (
	"testing"
)

func TestGoroutine(t *testing.T) {
	ch := make(chan string)
	go func() { ch <- "ping" }() // 没人收就会卡在这
	//msg := <-ch
	//fmt.Println(msg)
}
