package minstack

// Problem: Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
// Implement MinStack with the following operations:
//   Constructor() MinStack          -> initialize the data structure.
//   (MinStack).Push(val int)        -> push element val onto the stack.
//   (MinStack).Pop()                -> remove the element on the top of the stack.
//   (MinStack).Top() int            -> get the top element.
//   (MinStack).GetMin() int         -> retrieve the minimum element in the stack.
// The stack should always hold at least one element before Pop, Top, or GetMin is called.
//
// Example:
//   Input
//     ["MinStack","push","push","push","getMin","pop","top","getMin"]
//     [[],[-2],[0],[-3],[],[],[],[]]
//   Output
//     [null,null,null,null,-3,null,0,-2]
//
// Your task: Fill in the missing logic below so the provided tests pass without changing their expectations.

type MinStack struct{}

func Constructor() MinStack {
	panic("TODO: implement Constructor")
}

func (s *MinStack) Push(val int) {
	panic("TODO: implement Push")
}

func (s *MinStack) Pop() {
	panic("TODO: implement Pop")
}

func (s *MinStack) Top() int {
	panic("TODO: implement Top")
}

func (s *MinStack) GetMin() int {
	panic("TODO: implement GetMin")
}
