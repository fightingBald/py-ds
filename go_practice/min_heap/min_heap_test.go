package minheap

import "testing"

func TestMinHeapOrdering(t *testing.T) {
	h := New()

	inputs := []int{5, 1, 4, 2, 3}
	for _, v := range inputs {
		h.Push(v)
	}

	if got, ok := h.Peek(); !ok || got != 1 {
		t.Fatalf("peek expected (1, true), got (%d, %v)", got, ok)
	}

	var order []int
	for h.Len() > 0 {
		order = append(order, h.Pop())
	}

	want := []int{1, 2, 3, 4, 5}
	for i := range want {
		if order[i] != want[i] {
			t.Fatalf("pop #%d expected %d, got %d", i, want[i], order[i])
		}
	}
}

func TestMinHeapPeekEmpty(t *testing.T) {
	h := New()

	if _, ok := h.Peek(); ok {
		t.Fatalf("peek on empty should report ok=false")
	}
}

func TestMinHeapPopEmptyPanics(t *testing.T) {
	h := New()
	mustPanic(t, func() { h.Pop() })
}

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatalf("expected panic, got none")
		}
	}()
	fn()
}
