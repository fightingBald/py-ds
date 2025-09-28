package deque

// Problem: Implement a double-ended queue that mirrors Python's collections.deque behavior for integer values.
// Support the following operations, all in amortized O(1) time:
//   (d *Deque) Append(v int)      -> add v to the right end.
//   (d *Deque) AppendLeft(v int)  -> add v to the left end.
//   (d *Deque) Pop() int          -> remove and return the rightmost element; panic if empty.
//   (d *Deque) PopLeft() int      -> remove and return the leftmost element; panic if empty.
//   (d *Deque) Len() int          -> return the number of stored elements.
//
// Example sequence (mirrors Python):
//   var dq Deque
//   dq.Append(1)        // deque([1])
//   dq.AppendLeft(2)    // deque([2, 1])
//   dq.Append(3)        // deque([2, 1, 3])
//   dq.PopLeft() == 2   // deque([1, 3])
//   dq.Pop() == 3       // deque([1])
//
// Fill in the missing implementation so all provided tests pass without modifying their expectations.

type Deque struct{}

func (d *Deque) Append(v int) {
	panic("TODO: implement Append")
}

func (d *Deque) AppendLeft(v int) {
	panic("TODO: implement AppendLeft")
}

func (d *Deque) Pop() int {
	panic("TODO: implement Pop")
}

func (d *Deque) PopLeft() int {
	panic("TODO: implement PopLeft")
}

func (d *Deque) Len() int {
	panic("TODO: implement Len")
}
