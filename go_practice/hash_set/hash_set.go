package hashset

// Problem: Implement a hash set for strings that mirrors Python's built-in set behavior.
// Supported operations (the zero value of Set must be ready to use, just like Python's set()):
//   (s *Set) Add(v string)         -> add an element to the set (duplicates ignored).
//   (s *Set) Remove(v string)      -> remove an element; panic if the value is absent (matching Python's set.remove).
//   (s *Set) Contains(v string) bool -> report membership.
//   (s *Set) Len() int             -> number of stored elements.
//   (s *Set) Items() []string      -> return all elements in any order (mirrors set iteration order flexibility).
//
// Example usage similar to Python:
//   var s Set              // zero value is ready
//   s.Add("go")
//   s.Add("python")
//   s.Remove("go")
//   s.Contains("python") == true
//
// Complete the implementation so the provided tests pass. Do not change the tests.

type Set struct {
	m map[string]struct{}
}

func (s *Set) Contains(v string) bool {
	if s.m == nil {
		return false
	}
	_, ok := s.m[v]
	return ok
}

func (s *Set) Add(v string) {
	if s.Contains(v) {
		return
	}
	return
}

func (s *Set) Remove(v string) {
	panic("TODO: implement Remove")
}

func (s *Set) Len() int {
	panic("TODO: implement Len")
}

func (s *Set) Items() []string {
	panic("TODO: implement Items")
}
