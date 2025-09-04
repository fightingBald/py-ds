package a_basic_type

import (
	"fmt"
	"testing"
)

// 1. 把函数存到一个变量里，以后通过变量来调用。
func square(n int) int {
	return n * n
}
func TestSquare(t *testing.T) {
	f := square
	if f(3) != 9 {
		t.Fatal("expected 9")
	}
}

// 2; 把函数当成参数传给另一个函数。
// 参数类型是 一个 (int)->int 函数，。
func doubleOperator(n int, f func(int) int) int {
	return f(f(n))
}

func TestDoubleOperator(t *testing.T) {
	if doubleOperator(3, square) != 81 {
		t.Fatal("expected 81")
	}
}

// 3.函数还可以作为返回值。
// Go 里的函数不仅能被传入，还能被返回。
// 下面这个 makeAdder 函数，返回一个 func(int) int函数
// 👉 这就是 闭包（closure）：函数里还能带着外层的变量base。
func makeAdder(base int) func(int) int {
	return func(x int) int {
		return base + x
	}
}

func TestMakeAdder(t *testing.T) {
	add5 := makeAdder(5)  // 返回一个函数
	fmt.Println(add5(3))  // 8
	fmt.Println(add5(10)) // 15
}
