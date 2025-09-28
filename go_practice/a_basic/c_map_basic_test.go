package a_basic

import (
	"fmt"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapNativeCommonOps(t *testing.T) {
	t.Run("INIT_初始化_三板斧", func(t *testing.T) {
		lang2rank := map[string]int{"go": 1, "python": 2}
		// 2) make：可选容量提示（不是硬上限，只是增长预估）
		mHint := make(map[string]int, 16) // 预估会放 16 个键，减少扩容次数

		require.Equal(t, 2, len(lang2rank))

		require.Equal(t, 0, len(mHint))

		// 3) var：nil map 作为“空”的惯用表示
		var mNil map[string]int
		require.Nil(t, mNil)
		require.Equal(t, 0, len(mNil))
		require.Equal(t, 0, mNil["absent"]) // 读不存在的键 → 值类型的零值
		require.Panics(t, func() { mNil["x"] = 1 }, "往 nil map 写会 panic")
	})

	t.Run("CRUD_增删改查_含ok惯用法", func(t *testing.T) {
		m := map[string]int{}
		// Create/Update
		m["go"] = 1
		m["go"] = m["go"] + 1 // 更新
		m["py"]++             // 不存在时零值+1 → 1

		require.Equal(t, 2, m["go"])
		require.Equal(t, 1, m["py"])

		// Read 带 ok 惯用法
		if v, ok := m["java"]; ok {
			t.Fatalf("不该有java: %v", v)
		}

		// Delete
		delete(m, "py") // 删不存在也不会 panic
		_, ok := m["py"]
		require.False(t, ok)
	})

	t.Run("遍历_range_顺序不保证", func(t *testing.T) {
		m := map[string]int{"go": 1, "python": 2, "java": 3}
		for k, v := range m {
			fmt.Printf("key=%s val=%d\n", k, v) // 顺序是“随机”的，别依赖
		}
	})

	t.Run("遍历时删除是安全的", func(t *testing.T) {
		m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
		for k := range m {
			if k%2 == 0 {
				delete(m, k) // 删除当前遍历的键是安全的
			}
		}
		require.Equal(t, 2, len(m))
		_, ok2 := m[2]
		require.False(t, ok2)
	})

	t.Run("Map作为集合_Set模式", func(t *testing.T) {
		set := map[string]struct{}{}
		set["go"] = struct{}{}
		set["python"] = struct{}{}
		_, hasGo := set["go"]
		require.True(t, hasGo)
		// Remove
		delete(set, "go")
		_, hasGo = set["go"]
		require.False(t, hasGo)
	})

	t.Run("频次统计_Counter", func(t *testing.T) {
		cnt := map[rune]int{}
		for _, r := range "googo" {
			cnt[r]++
		}
		require.Equal(t, 3, cnt['o'])
		require.Equal(t, 2, cnt['g'])
	})

	t.Run("嵌套map_先初始化子map再写入", func(t *testing.T) {
		m := map[string]map[string]int{}
		if _, ok := m["user"]; !ok {
			m["user"] = make(map[string]int)
		}
		m["user"]["age"] = 18
		require.Equal(t, 18, m["user"]["age"])
	})

	t.Run("map与slice组合_append惯用法", func(t *testing.T) {
		m := map[string][]int{}
		m["a"] = append(m["a"], 1) // 不用先判断，append会处理零值nil切片
		m["a"] = append(m["a"], 2, 3)
		require.Equal(t, []int{1, 2, 3}, m["a"])
	})

	t.Run("清空_clear", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		clear(m)
		require.Equal(t, 0, len(m))
	})

	t.Run("键类型与结构体键", func(t *testing.T) {
		// 键必须是“comparable”的类型：基础类型、指针、数组、以及其字段全可比的 struct
		type Key struct {
			ID   int
			Name string
		}
		m := map[Key]string{
			{ID: 1, Name: "A"}: "x",
			{ID: 1, Name: "B"}: "y",
		}
		require.Equal(t, "x", m[Key{ID: 1, Name: "A"}])
	})

	t.Run("键列表与排序_maps.Keys", func(t *testing.T) {
		m := map[string]int{"b": 2, "a": 1, "c": 3}

		keys := slices.Collect(maps.Keys(m)) // iter.Seq[string] -> []string

		slices.Sort(keys)
		require.Equal(t, []string{"a", "b", "c"}, keys)
	})
}

func TestMapsLibBasics_Equal_Clone_DeleteFunc(t *testing.T) {
	t.Run("maps.Equal_值相等_忽略内部顺序", func(t *testing.T) {
		a := map[string]int{"x": 1, "y": 2}
		b := map[string]int{"y": 2, "x": 1}
		require.True(t, maps.Equal(a, b), "键值对集合一致就算相等")
	})

	t.Run("maps.Equal_nil视为空", func(t *testing.T) {
		var nilMap map[string]int
		empty := map[string]int{}
		require.True(t, maps.Equal(nilMap, empty))
	})

	t.Run("maps.Clone_深克隆键值对", func(t *testing.T) {
		src := map[string]int{"a": 1}
		cl := maps.Clone(src)
		cl["a"] = 999
		require.Equal(t, 1, src["a"])
		require.Equal(t, 999, cl["a"])
	})

	t.Run("maps.DeleteFunc_按条件批删", func(t *testing.T) {
		m := map[string]int{"odd": 1,
			"even": 2,
			"odd2": 3}
		maps.DeleteFunc(m, valueIsOdd)
		require.Equal(t, map[string]int{"even": 2}, m)
	})

	t.Run("maps给slice去重", func(t *testing.T) {
		langs := []string{"go", "python", "java", "python", "java", "python", "java", "python", "java"}

		set := map[string]struct{}{}

		for _, lang := range langs {
			set[lang] = struct{}{}
		}

		uniqueLangs := slices.Collect(maps.Keys(set))

		fmt.Printf("%v\n", uniqueLangs)

	})

}
func valueIsOdd(_ string, v int) bool {
	return v%2 == 1
}
