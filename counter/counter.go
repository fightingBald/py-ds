package counter

type Counter map[string]int

func (c Counter) Add(key string) {
	c[key]++
}

func (c Counter) Count(key string) int {
	return c[key]
}
