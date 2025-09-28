package ordereddict

import "testing"

func TestOrderedDictOrderAndMove(t *testing.T) {
	od := New()

	od.Set("py", "Python")
	od.Set("go", "Go")
	od.Set("rs", "Rust")
	od.Set("py", "Py3") // update should not change order

	keys := od.Keys()
	wantKeys := []string{"py", "go", "rs"}
	if len(keys) != len(wantKeys) {
		t.Fatalf("expected %d keys, got %d", len(wantKeys), len(keys))
	}
	for i := range wantKeys {
		if keys[i] != wantKeys[i] {
			t.Fatalf("keys[%d] want %s, got %s", i, wantKeys[i], keys[i])
		}
	}

	if val, ok := od.Get("py"); !ok || val != "Py3" {
		t.Fatalf("expected Get py to return Py3, got %q (ok=%v)", val, ok)
	}

	od.MoveToEnd("py", true)
	keys = od.Keys()
	wantKeys = []string{"go", "rs", "py"}
	if len(keys) != len(wantKeys) {
		t.Fatalf("expected %d keys after move last, got %d", len(wantKeys), len(keys))
	}
	for i := range wantKeys {
		if keys[i] != wantKeys[i] {
			t.Fatalf("after MoveToEnd last keys[%d] want %s, got %s", i, wantKeys[i], keys[i])
		}
	}

	od.MoveToEnd("py", false)
	keys = od.Keys()
	wantKeys = []string{"py", "go", "rs"}
	if len(keys) != len(wantKeys) {
		t.Fatalf("expected %d keys after move front, got %d", len(wantKeys), len(keys))
	}
	for i := range wantKeys {
		if keys[i] != wantKeys[i] {
			t.Fatalf("after MoveToEnd front keys[%d] want %s, got %s", i, wantKeys[i], keys[i])
		}
	}

	items := od.Items()
	wantItems := []Item{
		{Key: "py", Value: "Py3"},
		{Key: "go", Value: "Go"},
		{Key: "rs", Value: "Rust"},
	}
	if len(items) != len(wantItems) {
		t.Fatalf("expected %d items, got %d", len(wantItems), len(items))
	}
	for i := range wantItems {
		if items[i] != wantItems[i] {
			t.Fatalf("items[%d] want %+v, got %+v", i, wantItems[i], items[i])
		}
	}

	od.Delete("go")
	keys = od.Keys()
	wantKeys = []string{"py", "rs"}
	if len(keys) != len(wantKeys) {
		t.Fatalf("expected %d keys after delete, got %d", len(wantKeys), len(keys))
	}
	for i := range wantKeys {
		if keys[i] != wantKeys[i] {
			t.Fatalf("after delete keys[%d] want %s, got %s", i, wantKeys[i], keys[i])
		}
	}

	if od.Len() != 2 {
		t.Fatalf("expected len 2, got %d", od.Len())
	}
}

func TestOrderedDictDeleteMissingPanics(t *testing.T) {
	od := New()
	mustPanic(t, func() { od.Delete("missing") })
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
