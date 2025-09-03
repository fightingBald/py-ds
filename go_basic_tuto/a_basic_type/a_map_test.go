package a_basic_type

import (
	"reflect"
	"sort"
	"testing"
)

// ---------- 实现部分（你要练的函数） ----------

// OrderKeysByCount 返回按计数降序、key 升序（同分时）的 key 列表。
// 仅保留计数 > 0 的条目；不得修改输入 map。
// nil/空 map 返回空切片。
func OrderKeysByCount(m map[string]int) []string {
	if len(m) == 0 {
		return []string{}
	}
	type pair struct {
		k   string
		cnt int
	}
	arr := make([]pair, 0, len(m))
	for k, v := range m {
		if v > 0 {
			arr = append(arr, pair{k: k, cnt: v})
		}
	}
	if len(arr) == 0 {
		return []string{}
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].cnt != arr[j].cnt {
			return arr[i].cnt > arr[j].cnt // count 降序
		}
		return arr[i].k < arr[j].k // 同分按 key 升序
	})
	out := make([]string, len(arr))
	for i, p := range arr {
		out[i] = p.k
	}
	return out
}

// ---------- 测试部分（表驱动 + 子测试） ----------

func TestOrderKeysByCount_Basic(t *testing.T) {
	in := map[string]int{
		"apple":  3,
		"banana": 1,
		"cherry": 2,
	}
	got := OrderKeysByCount(in)
	want := []string{"apple", "cherry", "banana"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("basic mismatch:\n got: %#v\nwant: %#v", got, want)
	}
	// 不应修改原 map
	if in["banana"] != 1 {
		t.Fatalf("input map mutated: banana=%d", in["banana"])
	}
}

func TestOrderKeysByCount_TiesAndNonPositive(t *testing.T) {
	in := map[string]int{
		"b": 2,
		"a": 2,  // 与 b 同分，a 应该排在 b 前
		"x": 0,  // 忽略
		"y": -5, // 忽略
		"c": 1,
	}
	got := OrderKeysByCount(in)
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ties/non-positive mismatch:\n got: %#v\nwant: %#v", got, want)
	}
}

func TestOrderKeysByCount_EmptyAndNil(t *testing.T) {
	var nilMap map[string]int
	cases := []map[string]int{
		{},
		nilMap,
	}
	for i, in := range cases {
		got := OrderKeysByCount(in)
		if len(got) != 0 {
			t.Fatalf("case %d: want empty slice, got %#v", i, got)
		}
	}
}

func TestOrderKeysByCount_BigKeysStableRule(t *testing.T) {
	in := map[string]int{
		"mango": 5,
		"melon": 5, // 与 mango 同分，"mango" < "melon"
		"pear":  3,
		"plum":  3, // 与 pear 同分，"pear" < "plum"
		"fig":   1,
		"grape": 2,
		"guava": 0, // 忽略
	}
	got := OrderKeysByCount(in)
	want := []string{"mango", "melon", "pear", "plum", "grape", "fig"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("stable rule mismatch:\n got: %#v\nwant: %#v", got, want)
	}
}
