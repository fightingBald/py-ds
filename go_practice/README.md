# Go Practice vs. Python Collections

Exercises under this folder mimic familiar Python built-ins and `collections` types so you can translate muscle memory into Go. Each package contains a scaffolded implementation with TODOs and a table-driven test suite—fill in the code until `go test ./go_practice/...` passes.

| Python type | Go package | Notes |
|-------------|-----------|-------|
| `collections.deque` | `deque` | Bidirectional queue with append/pop on both ends. |
| `heapq` | `min_heap` | Binary min-heap supporting push, pop, peek. |
| `set` | `hash_set` | String set whose zero value works like `set()`; add/remove/contains with panic on missing remove. |
| `collections.Counter` | `counter` | Multiset counting strings with `MostCommon`. |
| `collections.defaultdict` | `defaultdict` | Generic map with default factory, sorted `Items`. |
| `collections.OrderedDict` | `ordered_dict` | Insertion-ordered dict with `MoveToEnd`. |

Future additions can extend coverage to other Python conveniences—open an issue or add a new package mirroring the interface you want to practice.
