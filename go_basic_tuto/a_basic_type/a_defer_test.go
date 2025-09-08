package a_basic_type

import (
	"fmt"
	"testing"
)

/*
defer = “注册一个收尾动作，等函数死之前帮你执行”。
语义上相当于 Python 里的 with 上下文管理器，或者 try/finally 里的 finally 块

defer f.Close()  // 自动收尾，不会忘关文件
defer mu.Unlock() // 保证退出前一定解锁
defer tx.Rollback()  // 不管啥情况先准备回滚
*/
func TestDefer(t *testing.T) {
	fmt.Println("start")

	defer fmt.Println("clean up 1")
	defer fmt.Println("clean up 2")

	fmt.Println("middle")
}
