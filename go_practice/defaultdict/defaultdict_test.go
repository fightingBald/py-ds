package defaultdict

import (
	"reflect"
	"testing"
)

func TestDefaultDictIntFactory(t *testing.T) {
	dd := New[string, int](func() int { return 0 })

	if dd.Len() != 0 {
		t.Fatalf("expected empty dict, got len=%d", dd.Len())
	}

	if got := dd.Get("python"); got != 0 {
		t.Fatalf("expected factory default 0, got %d", got)
	}

	if dd.Len() != 1 {
		t.Fatalf("expected len 1 after access, got %d", dd.Len())
	}

	dd.Set("go", 5)
	if got := dd.Get("go"); got != 5 {
		t.Fatalf("expected go value 5, got %d", got)
	}

	items := dd.Items()
	want := []Entry[string, int]{{Key: "go", Value: 5}, {Key: "python", Value: 0}}
	if !reflect.DeepEqual(items, want) {
		t.Fatalf("items mismatch. want %v, got %v", want, items)
	}

	dd.Clear()
	if dd.Len() != 0 {
		t.Fatalf("expected len 0 after Clear, got %d", dd.Len())
	}
}

func TestDefaultDictSliceAggregates(t *testing.T) {
	dd := New[string, []string](func() []string { return nil })

	values := map[string][]string{
		"py": []string{"django", "flask"},
		"go": []string{"chi"},
	}

	for key, list := range values {
		for _, entry := range list {
			current := dd.Get(key)
			current = append(current, entry)
			dd.Set(key, current)
		}
	}

	items := dd.Items()
	want := []Entry[string, []string]{
		{Key: "go", Value: []string{"chi"}},
		{Key: "py", Value: []string{"django", "flask"}},
	}

	if !reflect.DeepEqual(items, want) {
		t.Fatalf("items mismatch. want %v, got %v", want, items)
	}
}
