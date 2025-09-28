package counter

import (
	"fmt"
	"slices"
)

// Problem: Build a Counter type mirroring Python's collections.Counter for strings.
// Required behavior:
//   FromSlice(items []string) *Counter -> create a counter from initial data.
//   (c *Counter) Update(items []string) -> add counts from the slice.
//   (c *Counter) Count(value string) int -> return how many times value appears.
//   (c *Counter) MostCommon(n int) []Item -> top n items sorted by count descending then lexicographically.
//   (c *Counter) Items() []Item -> return all items sorted like MostCommon(len(counter)).
//
// Example aligned with Python:
//   c := FromSlice([]string{"go", "py", "go"}) // Counter({'go': 2, 'py': 1})
//   c.Update([]string{"py", "rust"})
//   c.Count("py") == 2
//   c.MostCommon(2) == []Item{{"go", 2}, {"py", 2}}
//
// Fill in the struct and methods so the accompanying tests pass.

type Item struct {
	Value string
	Count int
}

// The zero value of Counter must be ready to use, just like Python's Counter().

type Counter struct {
	counts map[string]int
}

func NewCounter() *Counter {
	return &Counter{counts: make(map[string]int)}
}

func NewCounterFromSlice(elems []string) *Counter {
	c := NewCounter()
	for _, elem := range elems {
		c.counts[elem]++
	}
	return c
}

func (c *Counter) Update(items []string) {
	for _, elem := range items {
		c.counts[elem]++
	}
}

func (c *Counter) Count(value string) int {
	return c.counts[value]
}

func (c *Counter) MostCommon(n int) []Item {
	mostCommon := make([]Item, 0, len(c.counts))
	for k, v := range c.counts {
		mostCommon = append(mostCommon, Item{Value: k, Count: v})
	}
	slices.SortFunc(mostCommon, func(i, j Item) int {
		return j.Count - i.Count
	})
	fmt.Println(mostCommon)
	return mostCommon[:n]
}
