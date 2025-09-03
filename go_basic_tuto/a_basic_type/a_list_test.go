package a_basic_type

import (
	"reflect"
	"testing"
)

// LeetCode 风格的切片（slice）练习，面向 Go 初学者。
// 每个练习包含：简要题目、函数签名（请实现）以及对应的测试用例。
// 练习覆盖的知识点：
// 1. 切片的创建与遍历
// 2. append 与容量（cap）
// 3. 切片视图与底层数组共享
// 4. nil 与空切片的区别
// 5. 使用 copy 创建独立拷贝
// 6. 从切片中删除元素
// 7. 在切片中插入元素
// 8. 反转切片
// 请在下方实现带有 TODO 的函数，然后运行 `go test -v ./...`。

// 1. SumInts：计算整数切片中所有元素之和。
// func SumInts(nums []int) int
func SumInts(nums []int) int {

	panic("implement me")
}

func TestSumInts(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		exp  int
	}{
		{"empty", []int{}, 0},
		{"one", []int{5}, 5},
		{"many", []int{1, 2, 3, 4}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumInts(tt.in); got != tt.exp {
				t.Fatalf("SumInts(%v) = %d, want %d", tt.in, got, tt.exp)
			}
		})
	}
}

// 2. AppendAndReturnCap：将 vals 依次 append 到 s，返回最终的容量 cap(s)。
// 目的是体会 append 与容量变化。
// func AppendAndReturnCap(s []int, vals []int) int
func AppendAndReturnCap(s []int, vals []int) int {
	panic("implement me")
}

func TestAppendAndReturnCap(t *testing.T) {
	s := make([]int, 0, 2)
	capAfter := AppendAndReturnCap(s, []int{1, 2, 3, 4})
	if capAfter < 4 {
		t.Fatalf("expected capacity >=4, got %d", capAfter)
	}
}

// 3. ModifyViewSet：传入一个切片视图（slice view），修改视图中的某个索引的值，返回被修改的视图。
// 目的是观察底层数组共享导致原切片变化。
// func ModifyViewSet(view []int, idx int, val int) []int
func ModifyViewSet(view []int, idx int, val int) []int {
	panic("implement me")
}

func TestModifyViewSet(t *testing.T) {
	base := []int{10, 20, 30, 40}
	view := base[1:3] // 包含 20,30
	_ = ModifyViewSet(view, 0, 99)
	// 修改 view 后，base 的对应位置也应改变
	if base[1] != 99 {
		t.Fatalf("expected base[1] to be 99, got %d", base[1])
	}
}

// 4. IsEmpty：判断切片是否为空（len == 0）。关注 nil 切片与空切片的区别。
// func IsEmpty(s []int) bool
func IsEmpty(s []int) bool {
	panic("implement me")
}

func TestIsEmpty(t *testing.T) {
	var a []int  // nil
	b := []int{} // empty but non-nil
	if !IsEmpty(a) || !IsEmpty(b) {
		t.Fatalf("both nil and empty slice should be treated as empty")
	}
}

// 5. CopyFirstN：返回 src 的前 n 个元素的独立拷贝（使用 copy），但长度不超过 src。
// func CopyFirstN(src []int, n int) []int
func CopyFirstN(src []int, n int) []int {
	panic("implement me")
}

func TestCopyFirstN(t *testing.T) {
	src := []int{1, 2, 3, 4}
	c := CopyFirstN(src, 3)
	if !reflect.DeepEqual(c, []int{1, 2, 3}) {
		t.Fatalf("unexpected copy result: %v", c)
	}
	// 修改拷贝不应影响原切片
	c[0] = 99
	if src[0] == 99 {
		t.Fatalf("modifying copy should not change src")
	}
}

// 6. DeleteAt：从切片 s 中删除索引 idx 的元素并返回新的切片（保持元素顺序）。
// func DeleteAt(s []int, idx int) []int
func DeleteAt(s []int, idx int) []int {
	panic("implement me")
}

func TestDeleteAt(t *testing.T) {
	s := []int{1, 2, 3, 4}

	a := DeleteAt(s, 2) // 删除值 3 (索引 2)
	if !reflect.DeepEqual(a, []int{1, 2, 4}) {
		t.Fatalf("DeleteAt wrong: %v", a)
	}
}

// 7. InsertAt：在切片 s 的索引 idx 处插入 val，并返回新的切片。
// func InsertAt(s []int, idx int, val int) []int
func InsertAt(s []int, idx int, val int) []int {
	panic("implement me")
}

func TestInsertAt(t *testing.T) {
	s := []int{1, 2, 4}

	a := InsertAt(s, 2, 3) // 在索引 2 插入 3
	if !reflect.DeepEqual(a, []int{1, 2, 3, 4}) {
		t.Fatalf("InsertAt wrong: %v", a)
	}
}

// 8. Reverse：原地反转切片并返回同一切片引用。
// func Reverse(s []int) []int
func Reverse(s []int) []int {
	panic("implement me")
}

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4}
	Reverse(s)
	if !reflect.DeepEqual(s, []int{4, 3, 2, 1}) {
		t.Fatalf("Reverse wrong: %v", s)
	}
}
