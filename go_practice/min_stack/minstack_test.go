package minstack

import "testing"

func TestMinStackExample1(t *testing.T) {
	stack := Constructor()

	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if got := stack.GetMin(); got != -3 {
		t.Fatalf("expected min -3 after pushes, got %d", got)
	}

	stack.Pop()

	if got := stack.Top(); got != 0 {
		t.Fatalf("expected top 0 after pop, got %d", got)
	}

	if got := stack.GetMin(); got != -2 {
		t.Fatalf("expected min -2 after pop, got %d", got)
	}
}

func TestMinStackInterleaved(t *testing.T) {
	stack := Constructor()

	type step struct {
		op       string
		value    int
		expected int
	}

	steps := []step{
		{op: "push", value: 5},
		{op: "push", value: 1},
		{op: "push", value: 5},
		{op: "getMin", expected: 1},
		{op: "pop"},
		{op: "top", expected: 1},
		{op: "getMin", expected: 1},
		{op: "pop"},
		{op: "getMin", expected: 5},
	}

	for i, st := range steps {
		switch st.op {
		case "push":
			stack.Push(st.value)
		case "pop":
			stack.Pop()
		case "top":
			if got := stack.Top(); got != st.expected {
				t.Fatalf("step %d (%s): expected %d, got %d", i, st.op, st.expected, got)
			}
		case "getMin":
			if got := stack.GetMin(); got != st.expected {
				t.Fatalf("step %d (%s): expected %d, got %d", i, st.op, st.expected, got)
			}
		default:
			t.Fatalf("unknown operation %q", st.op)
		}
	}
}
