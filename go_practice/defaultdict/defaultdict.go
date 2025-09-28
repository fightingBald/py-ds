package defaultdict

// Problem: Implement a generic DefaultDict similar to Python's collections.defaultdict.
// Behavior requirements:
//   New[K comparable, V any](factory func() V) *DefaultDict[K, V]
//   (d *DefaultDict[K, V]) Get(key K) V      -> returns the value for key, creating it with factory if missing.
//   (d *DefaultDict[K, V]) Set(key K, value V) -> assign value without calling the factory.
//   (d *DefaultDict[K, V]) Len() int         -> number of stored keys.
//   (d *DefaultDict[K, V]) Items() []Entry[K, V] -> entries sorted by key ascending for deterministic tests.
//   (d *DefaultDict[K, V]) Clear()           -> remove all keys.
//
// Example analogous to Python:
//   dd := New[string, []string](func() []string { return nil })
//   dd.Get("python")            // []string{}, and key now exists
//   dd.Set("go", []string{"tour"})
//   dd.Get("python") = append(...)
//
// Fill out the struct and methods so the tests compile and pass.

type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

type DefaultDict[K comparable, V any] struct{}

func New[K comparable, V any](factory func() V) *DefaultDict[K, V] {
	panic("TODO: implement New")
}

func (d *DefaultDict[K, V]) Get(key K) V {
	panic("TODO: implement Get")
}

func (d *DefaultDict[K, V]) Set(key K, value V) {
	panic("TODO: implement Set")
}

func (d *DefaultDict[K, V]) Len() int {
	panic("TODO: implement Len")
}

func (d *DefaultDict[K, V]) Items() []Entry[K, V] {
	panic("TODO: implement Items")
}

func (d *DefaultDict[K, V]) Clear() {
	panic("TODO: implement Clear")
}
