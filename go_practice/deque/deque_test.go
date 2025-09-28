package deque

import "testing"

type operation struct {
	name     string
	value    int
	expected int
}

func TestDequeSequence(t *testing.T) {
	dq := New()

	steps := []operation{
		{name: "append", value: 1},
		{name: "appendLeft", value: 2},
		{name: "append", value: 3},
		{name: "len", expected: 3},
		{name: "popLeft", expected: 2},
		{name: "pop", expected: 3},
		{name: "len", expected: 1},
		{name: "appendLeft", value: 5},
		{name: "pop", expected: 1},
		{name: "popLeft", expected: 5},
	}

	for i, step := range steps {
		switch step.name {
		case "append":
			dq.Append(step.value)
		case "appendLeft":
			dq.AppendLeft(step.value)
		case "pop":
			got := dq.Pop()
			if got != step.expected {
				t.Fatalf("step %d pop: want %d, got %d", i, step.expected, got)
			}
		case "popLeft":
			got := dq.PopLeft()
			if got != step.expected {
				t.Fatalf("step %d popLeft: want %d, got %d", i, step.expected, got)
			}
		case "len":
			if got := dq.Len(); got != step.expected {
				t.Fatalf("step %d len: want %d, got %d", i, step.expected, got)
			}
		default:
			t.Fatalf("unknown operation %q", step.name)
		}
	}
}

func TestDequePopEmptyPanics(t *testing.T) {
	dq := New()

	mustPanic(t, func() { dq.Pop() })
	mustPanic(t, func() { dq.PopLeft() })
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
