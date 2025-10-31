package deque

import "slices"

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

type Deque struct {
	queue []int
}

func (d *Deque) AppendLeft(v int) {
	d.queue = slices.Insert(d.queue, 0, v)
}

func (d *Deque) Append(v int) {
	d.queue = append(d.queue, v)
}

func (d *Deque) PopLeft() int {
	if len(d.queue) == 0 {
		panic("empty deque")
	}
	v := d.queue[0]
	d.queue = d.queue[1:]
	return v
}

func (d *Deque) Pop() int {
	if len(d.queue) == 0 {
		panic("empty deque")
	}
	v := d.queue[len(d.queue)-1]
	d.queue = d.queue[:len(d.queue)-1]
	return v
}

func (d *Deque) Len() int {
	return len(d.queue)
}
