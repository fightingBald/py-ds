package a_basic

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// 工具：保序删除 / 无序删除（O(1)）
func deleteOrdered[T any](s []T, i int) []T {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
func deleteUnordered[T any](s []T, i int) []T {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func TestSliceNativeCommonOps(t *testing.T) {
	t.Run("INIT_初始化_三板斧", func(t *testing.T) {
		// 字面量
		nums := []int{1, 2, 3}
		langs := []string{"go", "python", "java"}
		fmt.Printf("nums=%v langs=%v\n", nums, langs)

		// make：长度与容量
		s1 := make([]int, 3)     // len=3,cap=3，全零值
		s2 := make([]int, 0, 10) // len=0,cap=10，预留容量方便 append
		fmt.Printf("s1=%v (len=%d,cap=%d)\n", s1, len(s1), cap(s1))
		fmt.Printf("s2=%v (len=%d,cap=%d)\n", s2, len(s2), cap(s2))

		// var：nil 切片，常见空值表示
		var sNil []int
		require.Nil(t, sNil)
		require.Equal(t, 0, len(sNil))
		require.Equal(t, 0, cap(sNil))
	})
	t.Run("遍历", func(t *testing.T) {
		langs := []string{"go", "python", "java"}
		for i, lang := range langs {
			fmt.Printf("idx %d lang=%v\n", i, lang)
		}
	})
	t.Run("逆序遍历", func(t *testing.T) {
		langs := []string{"go", "python", "java"}
		for i := len(langs) - 1; i >= 0; i-- {
			fmt.Printf("idx %d lang=%v\n", i, langs[i])
		}
	})
	t.Run("原生支持_切片视图，注意共用底层数组陷阱", func(t *testing.T) {
		s := []int{10, 20, 30, 40}
		sView := s[1:]
		//对底层数组的一段“窗口视图” 修改视图里的元素，就是直接改底层数组 → 所有指向同一数组的切片都会跟着变。
		//注意：append 会“断开关系”
		//如果你在子切片上 append，且 没超过 cap → 仍然是原数组，修改会影响原切片。
		//如果 append 触发扩容 → Go 会新建数组，子切片就和原切片“断亲了”。
		sView[0] = 8888
		fmt.Println(sView)
		fmt.Println(s) //[10 20 80 40]
		sView = append(sView, 9999)
		fmt.Println(sView)
		fmt.Println(s) //[10 20 80 40]

	})
	t.Run("原生支持_append", func(t *testing.T) {
		s := make([]int, 3, 5)
		s = append(s, 1, 2, 3, 4, 5)
		fmt.Println(s) //[0 0 0 1 2 3 4 5]

		// 批量追加
		more := []int{999, 999}
		//... 叫展开操作符
		//如果你忘了写 ...，比如 append(s, more)，那是往 s 里塞一个元素，这个元素的类型还是 []int，直接报编译错误。
		s = append(s, more...)
		fmt.Println(s) //[0 0 0 1 2 3 4 5 999 999]

	})
	t.Run("原生支持_copy，底层数组隔离", func(t *testing.T) {
		//copy(dst, src) 会把 src 里的元素往 dst 里搬，能搬多少搬多少，按下标从 0 开始一一覆盖。
		src := []int{1, 2, 3, 4}
		dst := make([]int, 2) //[0,0]

		n := copy(dst, src)        // 复制 src 前2个元素
		fmt.Println("dst =", dst)  // [1 2]
		fmt.Println("copied =", n) // 2

		// 真正完整拷贝
		dst2 := make([]int, len(src))
		copy(dst2, src)
		fmt.Println("dst2 =", dst2) // [1 2 3 4]

	})
	t.Run("翻转reverse", func(t *testing.T) {
		s := []int{10, 20, 30, 40}
		slices.Reverse(s)
		fmt.Println(s)
	})
}
func TestSlicesLibBasics(t *testing.T) {
	t.Run("80/20 常用清单", func(t *testing.T) {
		s := []int{10, 20, 30, 40}
		_ = slices.Contains(s, 20)              // 查有没有
		_ = slices.Index(s, 30)                 // 找位置
		_ = slices.Clone(s)                     // 深拷贝
		_ = slices.Concat(s, make([]int, 3, 5)) // 拼接
		_ = slices.Insert(s, 1, 15, 16)         // 在 1 插入 -> [10,15,16,20,30,40]
		_ = slices.Delete(s, 2, 4)              // 删 [2:4) -> [10,15,30,40]
		_ = slices.Replace(s, 1, 3, 11, 12)     // 先删掉区间 [i:j)，再在位置 i 插入 v...。

		slices.Sort(s)         // 排序
		slices.Reverse(s)      // 翻转
		_ = slices.IsSorted(s) // 是否已排好

		min := slices.Min([]int{3, 1, 9}) // 极值
		max := slices.Max([]int{3, 1, 9})

		s = slices.Grow(s, 100) // 预留容量
		s = slices.Clip(s[:2])  // 收缩 cap 到 len
		_ = min
		_ = max

	})
	t.Run("按照key来sort SortFunc", func(t *testing.T) {
		langs := []string{"go", "python", "java"}
		slices.SortFunc(langs, sortByLen)

		fmt.Printf("%v\n", langs)
	})
	t.Run("比较两个slice值是否相等slices.Equal", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{1, 2, 3}
		_ = slices.Equal(a, b) // true
	})

	t.Run("Clone_Equal_独立内存", func(t *testing.T) {
		orig := []string{"go", "python", "java"}
		clone := slices.Clone(orig) // 真正复制到底层新数组
		require.True(t, slices.Equal(orig, clone))
		require.NotSame(t, &orig[0], &clone[0]) // 地址不同，证明不共享底层
	})

	t.Run("Contains_Index_查找基本款", func(t *testing.T) {
		ss := []string{"go", "python", "java", "go"}
		require.True(t, slices.Contains(ss, "go"))
		require.Equal(t, 0, slices.Index(ss, "go")) // 只返第一个命中位置
	})

	t.Run("Concat_拼接优雅替代_append链", func(t *testing.T) {
		a := []int{1, 2}
		b := []int{3}
		c := []int{4, 5}
		got := slices.Concat(a, b, c)
		require.Equal(t, []int{1, 2, 3, 4, 5}, got)
	})

	t.Run("Insert_Delete_Replace_增删改", func(t *testing.T) {
		s := []int{10, 20, 30, 40}

		// Insert：在索引 1 插入多个元素
		s = slices.Insert(s, 1, 15, 16)
		require.Equal(t, []int{10, 15, 16, 20, 30, 40}, s)

		// Delete：[i:j) 半开区间删除
		s = slices.Delete(s, 2, 4) // 删除 16,20
		require.Equal(t, []int{10, 15, 30, 40}, s)

		// Replace：把一段替换为另一批
		s = slices.Replace(s, 1, 3, 11, 12, 13) // 15,30 → 11,12,13
		require.Equal(t, []int{10, 11, 12, 13, 40}, s)
	})

	t.Run("Sort_SortFunc_Reverse_排序与翻转", func(t *testing.T) {
		langs := []string{"Go", "python", "Java", "rust"}
		slices.Sort(langs) // 默认字典序（区分大小写）
		require.True(t, slices.IsSorted(langs))

		// 忽略大小写排序：SortFunc + cmp.Compare
		langs2 := []string{"Go", "python", "Java", "rust"}
		slices.SortFunc(langs2, func(a, b string) int {
			return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
		})
		require.True(t, slices.IsSortedFunc(langs2, func(a, b string) int {
			return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
		}))

		// Reverse：原地翻转
		nums := []int{1, 2, 3, 4}
		slices.Reverse(nums)
		require.Equal(t, []int{4, 3, 2, 1}, nums)
	})

	t.Run("Min_Max_极值", func(t *testing.T) {
		require.Equal(t, 1, slices.Min([]int{3, 1, 9}))
		require.Equal(t, 9, slices.Max([]int{3, 1, 9}))
	})

}

func sortByLen(a, b string) int {
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}
	return 0
}
