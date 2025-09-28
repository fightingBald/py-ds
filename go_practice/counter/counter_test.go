package counter

import "testing"

func TestCounterCountsAndMostCommon(t *testing.T) {
	c := FromSlice([]string{"python", "go", "python", "rust"})
	c.Update([]string{"go", "python", "go"})

	if got := c.Count("python"); got != 3 {
		t.Fatalf("expected python count 3, got %d", got)
	}
	if got := c.Count("go"); got != 3 {
		t.Fatalf("expected go count 3, got %d", got)
	}
	if got := c.Count("rust"); got != 1 {
		t.Fatalf("expected rust count 1, got %d", got)
	}
	if got := c.Count("java"); got != 0 {
		t.Fatalf("expected java count 0, got %d", got)
	}

	top := c.MostCommon(2)
	wantTop := []Item{{Value: "go", Count: 3}, {Value: "python", Count: 3}}
	for i := range wantTop {
		if top[i] != wantTop[i] {
			t.Fatalf("most common mismatch at %d: want %+v, got %+v", i, wantTop[i], top[i])
		}
	}

	all := c.Items()
	wantAll := []Item{{"go", 3}, {"python", 3}, {"rust", 1}}
	for i := range wantAll {
		if all[i] != wantAll[i] {
			t.Fatalf("items[%d] mismatch: want %+v, got %+v", i, wantAll[i], all[i])
		}
	}
}

func TestCounterMostCommonMoreThanSize(t *testing.T) {
	var c Counter
	c.Update([]string{"a", "b", "a"})

	got := c.MostCommon(10)
	want := []Item{{"a", 2}, {"b", 1}}

	if len(got) != len(want) {
		t.Fatalf("expected %d results, got %d", len(want), len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("mostCommon[%d] want %+v, got %+v", i, want[i], got[i])
		}
	}
}
