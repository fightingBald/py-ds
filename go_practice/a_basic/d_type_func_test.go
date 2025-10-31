package a_basic

import (
	"fmt"
	"testing"
)

// 函数“有类型”，比如 func(int) int，这就是个类型签名。
func TestFuncType(t *testing.T) {
	f := func(a int) int { return a * 2 }
	fmt.Println(f(3)) // 6
}

type Adder func(int, int) int

// Adder 是一个类型，它的值必须是“接受两个 int，返回一个 int 的函数”。
func apply(op Adder, x, y int) int {
	return op(x, y)
}

func TestFuncType1(t *testing.T) {
	sum := func(a, b int) int { return a + b }
	fmt.Println(apply(sum, 2, 3)) // 5
}
