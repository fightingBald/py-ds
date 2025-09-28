package counter

// Problem: Build a Counter type mirroring Python's collections.Counter for strings.
// Required behavior:
//   New() *Counter                    -> create an empty counter.
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

type Counter struct{}

func New() *Counter {
	panic("TODO: implement New")
}

func FromSlice(items []string) *Counter {
	panic("TODO: implement FromSlice")
}

func (c *Counter) Update(items []string) {
	panic("TODO: implement Update")
}

func (c *Counter) Count(value string) int {
	panic("TODO: implement Count")
}

func (c *Counter) MostCommon(n int) []Item {
	panic("TODO: implement MostCommon")
}

func (c *Counter) Items() []Item {
	panic("TODO: implement Items")
}
