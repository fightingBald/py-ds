package counter

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// Counter is a Python-like counter: maps keys to int64 counts.
// Zero value is ready to use (just like a map's zero value).
type Counter[T comparable] map[T]int64

// New creates an empty Counter.
func New[T comparable]() Counter[T] { return make(Counter[T]) }

// FromSlice builds a Counter from a slice in one shot.
func FromSlice[T comparable](xs []T) Counter[T] {
	c := New[T]()
	c.UpdateSlice(xs)
	return c
}

// Clone returns a deep copy of the Counter.
func (c Counter[T]) Clone() Counter[T] {
	out := make(Counter[T], len(c))
	for k, v := range c {
		out[k] = v
	}
	return out
}

// Get returns count for key (0 if missing).
func (c Counter[T]) Get(k T) int64 { return c[k] }

// Set sets the count for key (<=0 allowed; call Clean to drop zeros).
func (c Counter[T]) Set(k T, n int64) { c[k] = n }

// Add increments key by 1.
func (c Counter[T]) Add(k T) { c[k]++ }

// AddN increments key by n (n can be negative).
// If the resulting count is 0, the key is removed to keep map tidy.
func (c Counter[T]) AddN(k T, n int64) {
	if n == 0 {
		return
	}
	c[k] += n
	if c[k] == 0 {
		delete(c, k)
	}
}

// Update adds counts from another Counter (like Python's update()).
func (c Counter[T]) Update(other Counter[T]) {
	for k, v := range other {
		if v != 0 {
			c[k] += v
			if c[k] == 0 {
				delete(c, k)
			}
		}
	}
}

// UpdateSlice counts items in the slice and adds them in-place.
func (c Counter[T]) UpdateSlice(xs []T) {
	for _, x := range xs {
		c[x]++
	}
}

// Subtract subtracts counts from another Counter (like Python's subtract()).
func (c Counter[T]) Subtract(other Counter[T]) {
	for k, v := range other {
		if v != 0 {
			c[k] -= v
			if c[k] == 0 {
				delete(c, k)
			}
		}
	}
}

// Delete removes a key and returns its previous count (0 if not present).
func (c Counter[T]) Delete(k T) (prev int64) {
	prev = c[k]
	delete(c, k)
	return
}

// Clean deletes entries whose count == 0.
func (c Counter[T]) Clean() {
	for k, v := range c {
		if v == 0 {
			delete(c, k)
		}
	}
}

// Total returns the sum of all positive counts (matches common Python usage).
// If you need raw sum (including negatives), use TotalRaw.
func (c Counter[T]) Total() int64 {
	var s int64
	for _, v := range c {
		if v > 0 {
			s += v
		}
	}
	return s
}

// TotalRaw returns the algebraic sum of all counts (can be negative).
func (c Counter[T]) TotalRaw() int64 {
	var s int64
	for _, v := range c {
		s += v
	}
	return s
}

// Keys returns all keys (order undefined).
func (c Counter[T]) Keys() []T {
	out := make([]T, 0, len(c))
	for k := range c {
		out = append(out, k)
	}
	return out
}

// Values returns all counts (order undefined).
func (c Counter[T]) Values() []int64 {
	out := make([]int64, 0, len(c))
	for _, v := range c {
		out = append(out, v)
	}
	return out
}

// Pair holds a (Key, Count) entry.
type Pair[T comparable] struct {
	Key   T
	Count int64
}

// Items returns all (key, count) pairs (unsorted).
func (c Counter[T]) Items() []Pair[T] {
	out := make([]Pair[T], 0, len(c))
	for k, v := range c {
		out = append(out, Pair[T]{k, v})
	}
	return out
}

// MostCommon returns top-n pairs sorted by Count desc (then by key asc for stability).
// If n <= 0 or n >= len(c), returns all sorted.
func (c Counter[T]) MostCommon(n int) []Pair[T] {
	items := c.Items()
	sort.Slice(items, func(i, j int) bool {
		if items[i].Count == items[j].Count {
			// best effort: for comparable T we try fmt.Sprint for stable tie-break
			return fmt.Sprint(items[i].Key) < fmt.Sprint(items[j].Key)
		}
		return items[i].Count > items[j].Count
	})
	if n <= 0 || n >= len(items) {
		return items
	}
	return items[:n]
}

// Elements returns a slice that repeats each element according to its count (like Python's elements()).
// Negative or zero counts are ignored.
func (c Counter[T]) Elements() []T {
	total := 0
	for _, v := range c {
		if v > 0 {
			total += int(v)
		}
	}
	out := make([]T, 0, total)
	for k, v := range c {
		for i := int64(0); i < v; i++ {
			out = append(out, k)
		}
	}
	return out
}

// String implements fmt.Stringer for readable debug print, e.g. Counter{'a': 2, 'b': 1}
func (c Counter[T]) String() string {
	var b strings.Builder
	b.WriteString("Counter{")
	first := true
	for k, v := range c {
		if !first {
			b.WriteString(", ")
		}
		first = false
		fmt.Fprintf(&b, "%v:%d", k, v)
	}
	b.WriteString("}")
	return b.String()
}

//// Thread-safe variant ////////////////////////////////////////////////////////

// SafeCounter wraps Counter with a RWMutex for concurrent use.
// Rule of thumb: use SafeCounter only if multiple goroutines WRITE concurrently.
type SafeCounter[T comparable] struct {
	mu sync.RWMutex
	c  Counter[T]
}

// NewSafe creates a SafeCounter.
func NewSafe[T comparable]() *SafeCounter[T] { return &SafeCounter[T]{c: New[T]()} }

// Unsafe exposes the underlying Counter for advanced cases (use with care).
func (s *SafeCounter[T]) Unsafe() Counter[T] { return s.c }

// The methods below mirror Counter's API with locking.

func (s *SafeCounter[T]) Get(k T) int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.c[k]
}
func (s *SafeCounter[T]) Set(k T, n int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.c.Set(k, n)
}
func (s *SafeCounter[T]) Add(k T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.c.Add(k)
}
func (s *SafeCounter[T]) AddN(k T, n int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.c.AddN(k, n)
}
func (s *SafeCounter[T]) Update(other Counter[T]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.c.Update(other)
}
func (s *SafeCounter[T]) UpdateSlice(xs []T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.c.UpdateSlice(xs)
}
func (s *SafeCounter[T]) Subtract(other Counter[T]) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.c.Subtract(other)
}
func (s *SafeCounter[T]) Delete(k T) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.c.Delete(k)
}
func (s *SafeCounter[T]) Clean() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.c.Clean()
}
func (s *SafeCounter[T]) Total() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.c.Total()
}
func (s *SafeCounter[T]) TotalRaw() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.c.TotalRaw()
}
func (s *SafeCounter[T]) Keys() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.c.Keys()
}
func (s *SafeCounter[T]) Values() []int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.c.Values()
}
func (s *SafeCounter[T]) Items() []Pair[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.c.Items()
}
func (s *SafeCounter[T]) MostCommon(n int) []Pair[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.c.MostCommon(n)
}
func (s *SafeCounter[T]) Elements() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.c.Elements()
}
func (s *SafeCounter[T]) Clone() Counter[T] {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.c.Clone()
}
