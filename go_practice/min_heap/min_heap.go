package minheap

// Problem: Implement a binary min-heap that mimics Python's heapq interface for integers.
// Support the following operations:
//   (h *MinHeap) Push(v int)  -> push a value onto the heap.
//   (h *MinHeap) Pop() int    -> pop and return the smallest item; panic if the heap is empty.
//   (h *MinHeap) Peek() (int, bool) -> return the smallest item without removing it. The bool is false if empty.
//   (h *MinHeap) Len() int    -> number of stored elements.
// Ensure the operations run in O(log n) time for push/pop and O(1) for peek/len.
//
// Example sequence matching heapq:
//   var h MinHeap
//   h.Push(3)
//   h.Push(1)
//   h.Push(2)
//   h.Pop() == 1
//   h.Peek() == (2, true)
//
// Fill in the implementation so the provided tests succeed without altering their expectations.

type MinHeap struct{}

func (h *MinHeap) Push(v int) {
	panic("TODO: implement Push")
}

func (h *MinHeap) Pop() int {
	panic("TODO: implement Pop")
}

func (h *MinHeap) Peek() (int, bool) {
	panic("TODO: implement Peek")
}

func (h *MinHeap) Len() int {
	panic("TODO: implement Len")
}
