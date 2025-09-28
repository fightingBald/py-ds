package lrucache

import "testing"

type cacheOp struct {
	name     string
	key      int
	value    int
	expected int
}

func execute(t *testing.T, capacity int, ops []cacheOp, wants []int) {
	cache := Constructor(capacity)
	var got []int

	for i, op := range ops {
		switch op.name {
		case "put":
			cache.Put(op.key, op.value)
		case "get":
			value := cache.Get(op.key)
			got = append(got, value)
		default:
			t.Fatalf("unknown operation %q at step %d", op.name, i)
		}
	}

	if len(wants) != len(got) {
		t.Fatalf("expected %d get results, got %d", len(wants), len(got))
	}

	for i := range wants {
		if wants[i] != got[i] {
			t.Fatalf("get #%d: expected %d, got %d", i, wants[i], got[i])
		}
	}
}

func TestLRUCacheExampleSequence(t *testing.T) {
	ops := []cacheOp{
		{name: "put", key: 1, value: 1},
		{name: "put", key: 2, value: 2},
		{name: "get", key: 1},
		{name: "put", key: 3, value: 3},
		{name: "get", key: 2},
		{name: "put", key: 4, value: 4},
		{name: "get", key: 1},
		{name: "get", key: 3},
		{name: "get", key: 4},
	}

	wants := []int{1, -1, -1, 3, 4}

	execute(t, 2, ops, wants)
}

func TestLRUCacheUpdatesRefreshRecency(t *testing.T) {
	ops := []cacheOp{
		{name: "put", key: 2, value: 1},
		{name: "put", key: 1, value: 1},
		{name: "put", key: 2, value: 3},
		{name: "put", key: 4, value: 1},
		{name: "get", key: 1},
		{name: "get", key: 2},
	}

	wants := []int{-1, 3}

	execute(t, 2, ops, wants)
}
