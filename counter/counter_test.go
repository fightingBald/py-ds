// counter/counter_test.go
package counter_test

import (
	"github.com/fightingBald/py-ds/counter"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\n got: %#v\nwant: %#v", got, want)
	}
}

func assertMapEq[K comparable](t *testing.T, got counter.Counter[K], want map[K]int64) {
	t.Helper()
	// 复制并清理零值键，避免实现细节差异
	gc := got.Clone()
	gc.Clean()
	if len(gc) != len(want) {
		t.Fatalf("len mismatch: got %d want %d; got=%v", len(gc), len(want), gc)
	}
	for k, v := range want {
		if gc.Get(k) != v {
			t.Fatalf("key %v: got %d want %d; full=%v", k, gc.Get(k), v, gc)
		}
	}
}

func TestAddSetClean(t *testing.T) {
	c := counter.New[string]()
	c.Add("x")
	c.AddN("y", 3)
	c.AddN("y", -3) // 回到 0 应删除
	c.Set("z", 0)   // 手动放 0 ，Clean 再删
	c.Clean()
	assertMapEq(t, c, map[string]int64{"x": 1})
}

func TestUpdateAndSubtract(t *testing.T) {
	c := counter.FromSlice([]string{"a", "b", "a"})
	other := counter.FromSlice([]string{"a", "c"})
	c.Update(other) // a:3 b:1 c:1
	assertMapEq(t, c, map[string]int64{"a": 3, "b": 1, "c": 1})

	c.Subtract(counter.FromSlice([]string{"a", "a", "c"})) // a:1 b:1
	assertMapEq(t, c, map[string]int64{"a": 1, "b": 1})
}

func TestCloneIndependence(t *testing.T) {
	c := counter.FromSlice([]string{"a", "a", "b"})
	clone := c.Clone()
	clone.Add("a")
	assertEqual(t, c.Get("a"), int64(2))
	assertEqual(t, clone.Get("a"), int64(3))
}

func TestUpdateSliceLarge(t *testing.T) {
	text := strings.Repeat("go go go! ", 1000)
	words := strings.FieldsFunc(text, func(r rune) bool { return r == ' ' || r == '!' })
	c := counter.New[string]()
	c.UpdateSlice(words)
	// words 里有 ["go","go","go",""] 等，空串会被过滤（FieldsFunc不会产生空串）
	assertEqual(t, c.Get("go"), int64(3000))
}

func TestSafeCounterRaceFree(t *testing.T) {
	sc := counter.NewSafe[int]()
	const N = 1000
	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(v int) {
			defer wg.Done()
			sc.Add(v % 10)
			time.Sleep(time.Microsecond)
		}(i)
	}
	wg.Wait()
	// 每个 0..9 大约加了 N/10 次，确切值不要求；这里验证总和
	assertEqual(t, sc.Total(), int64(N))
}

func TestTableDriven(t *testing.T) {
	type op func(c counter.Counter[string])
	cases := []struct {
		name string
		ops  []op
		want map[string]int64
	}{
		{
			name: "basic add/subtract",
			ops: []op{
				func(c counter.Counter[string]) { c.Add("a") },
				func(c counter.Counter[string]) { c.AddN("a", 2) },
				func(c counter.Counter[string]) { c.AddN("b", 1) },
				func(c counter.Counter[string]) {
					c.Subtract(counter.Counter[string]{"a": 1})
				},
			},
			want: map[string]int64{"a": 2, "b": 1},
		},
		{
			name: "update from other",
			ops: []op{
				func(c counter.Counter[string]) {
					c.Update(counter.Counter[string]{"x": 2, "y": -2})
				},
				func(c counter.Counter[string]) { c.Clean() },
			},
			want: map[string]int64{"x": 2}, // y 被清理掉
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := counter.New[string]()
			for _, f := range tc.ops {
				f(c)
			}
			assertMapEq(t, c, tc.want)
		})
	}
}

// -------------------- Fuzz（Go 1.18+） --------------------

// 不变式：Elements() 长度 == Total()（只统计正数）
func FuzzElementsMatchesTotal(f *testing.F) {
	f.Add("aaabbbccc") // 种子
	f.Add("hello world")
	f.Add("")
	f.Fuzz(func(t *testing.T, s string) {
		c := counter.New[rune]()
		for _, r := range s {
			c.Add(r)
		}
		if int64(len(c.Elements())) != c.Total() {
			t.Fatalf("elements len %d != total %d; c=%v", len(c.Elements()), c.Total(), c)
		}
	})
}

// -------------------- Benchmark --------------------

func BenchmarkUpdateSlice_Small(b *testing.B) {
	words := []string{"a", "b", "a", "c", "a", "b"}
	for i := 0; i < b.N; i++ {
		c := counter.New[string]()
		c.UpdateSlice(words)
	}
}

func BenchmarkMostCommon_Top10(b *testing.B) {
	// 构造 10k 键，Zipf 分布
	c := counter.New[int]()
	for i := 0; i < 10000; i++ {
		c.AddN(i%1000, 1)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = c.MostCommon(10)
	}
}

// 也可以比较 SafeCounter 的开销
func BenchmarkSafeCounter_Add(b *testing.B) {
	sc := counter.NewSafe[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sc.Add(i & 1023)
	}
}
