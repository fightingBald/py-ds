package lrucache

// Problem: Design a data structure that supports the Least Recently Used (LRU) cache mechanism.
// Implement LRUCache with the following operations, all running in average O(1) time:
//   Constructor(capacity int) LRUCache -> instantiate the cache with a positive capacity.
//   (LRUCache).Get(key int) int        -> return the value of the key if it exists, otherwise -1.
//   (LRUCache).Put(key int, value int) -> update the value of the key if it exists or insert a new key-value pair.
//                                       When the cache reaches capacity, it should invalidate the least recently used entry.
//
// Example:
//   Input
//     ["LRUCache","put","put","get","put","get","put","get","get","get"]
//     [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
//   Output
//     [null,null,null,1,null,-1,null,-1,3,4]
//
// Fill in the cache implementation so the provided tests pass without modifying their expectations.

type LRUCache struct{}

func Constructor(capacity int) LRUCache {
	panic("TODO: implement Constructor")
}

func (c *LRUCache) Get(key int) int {
	panic("TODO: implement Get")
}

func (c *LRUCache) Put(key int, value int) {
	panic("TODO: implement Put")
}
