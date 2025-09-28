package a_basic

import (
	"slices"
)

// ===================== 辅助函数：优雅操作 =====================

// MapSlice 把每个元素映射成新值（返回新切片）
func MapSlice[T any, R any](in []T, f func(T) R) []R {
	out := make([]R, 0, len(in))
	for _, v := range in {
		out = append(out, f(v))
	}
	return out
}

// FilterSlice 过滤元素（返回新切片）
func FilterSlice[T any](in []T, pred func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

// ReduceSlice 归约：把一堆元素揉成一个值
func ReduceSlice[T any, R any](in []T, init R, f func(R, T) R) R {
	acc := init
	for _, v := range in {
		acc = f(acc, v)
	}
	return acc
}

// Clone 更清晰：等价于 slices.Clone
func Clone[T any](in []T) []T { return slices.Clone(in) }

// DeleteAt 有序删除：保持顺序，O(n)
func DeleteAt[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		panic("index out of range")
	}
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

// DeleteAtUnordered 无序删除：不保证顺序，O(1)
func DeleteAtUnordered[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		panic("index out of range")
	}
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// InsertAt 在 i 处插入一个元素，返回新切片（可能扩容）
func InsertAt[T any](s []T, i int, v T) []T {
	if i < 0 || i > len(s) {
		panic("index out of range")
	}
	s = append(s, v)     // 先把长度+1
	copy(s[i+1:], s[i:]) // i 之后整体右移
	s[i] = v
	return s
}

// Unique 保留首次出现，去重（稳定）
func Unique[T comparable](s []T) []T {
	seen := make(map[T]struct{}, len(s))
	out := make([]T, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

// Chunk 把切片按固定块大小切分
func Chunk[T any](s []T, size int) [][]T {
	if size <= 0 {
		panic("size must be > 0")
	}
	var res [][]T
	for i := 0; i < len(s); i += size {
		end := i + size
		if end > len(s) {
			end = len(s)
		}
		res = append(res, s[i:end])
	}
	return res
}
