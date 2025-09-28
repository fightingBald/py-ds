package ordereddict

// Problem: Implement an OrderedDict similar to Python's collections.OrderedDict for string keys and values.
// Support these operations (the zero value of OrderedDict must be ready to use):
//   (od *OrderedDict) Set(key string, value string) -> insert or update while preserving insertion order.
//   (od *OrderedDict) Get(key string) (string, bool) -> retrieve value, reporting if the key exists.
//   (od *OrderedDict) Delete(key string)            -> remove a key; panic if missing.
//   (od *OrderedDict) MoveToEnd(key string, last bool) -> move key to end if last=true, otherwise front.
//   (od *OrderedDict) Len() int
//   (od *OrderedDict) Keys() []string
//   (od *OrderedDict) Items() []Item -> ordered pairs representing the dictionary state.
//
// Example follow-up from Python docs:
//   var od OrderedDict
//   od.Set("py", "Python")
//   od.Set("go", "Go")
//   od.MoveToEnd("py", true) // py now last
//
// Implement the structure and methods so the tests pass.

type Item struct {
	Key   string
	Value string
}

type OrderedDict struct{}

func (od *OrderedDict) Set(key string, value string) {
	panic("TODO: implement Set")
}

func (od *OrderedDict) Get(key string) (string, bool) {
	panic("TODO: implement Get")
}

func (od *OrderedDict) Delete(key string) {
	panic("TODO: implement Delete")
}

func (od *OrderedDict) MoveToEnd(key string, last bool) {
	panic("TODO: implement MoveToEnd")
}

func (od *OrderedDict) Len() int {
	panic("TODO: implement Len")
}

func (od *OrderedDict) Keys() []string {
	panic("TODO: implement Keys")
}

func (od *OrderedDict) Items() []Item {
	panic("TODO: implement Items")
}
