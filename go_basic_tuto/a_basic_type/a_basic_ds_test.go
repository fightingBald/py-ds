package a_basic_type

import "testing"

func TestBasicArrSlice(t *testing.T) {
	//Arr 长度固定且不可变
	arr := [3]int{1, 2, 3}
	t.Logf("arr = %v lenth is %v", arr, len(arr))

	//Slice 长度可变
	slice := []int{1, 2, 3} //1.literal 创建
	t.Logf("slice = %v lenth is %v", slice, len(slice))
	slice = append(slice, 4, 5, 6)
	t.Logf("after append, slice = %v lenth is %v ", slice, len(slice))

	//2. make 创建
	slice2 := make([]int, 3, 5)   //slice2 = [0 0 0] lenth is 3 cap is 5
	slice2 = append(slice2, 4, 5) // slice2 = [0 0 0 4 5] lenth is 5 cap is 5

	t.Logf("slice2 = %v lenth is %v cap is %v", slice2, len(slice2), cap(slice2))
	t.Logf("after append, slice2 = %v lenth is %v cap is %v", slice2, len(slice2), cap(slice2))
}

func TestMap(t *testing.T) {
	// map 的创建
	m := map[string]int{
		"apple":  5,
		"banana": 3,
	}
	t.Logf("map m = %v", m) // map m = map[apple:5 banana:3]

	mByMake := make(map[string]int) // 使用 make 创建空 map
	mByMake["orange"] = 4

	// 添加或更新键值对
	m["cherry"] = 7 // 添加新键值对
	m["banana"] = 4 // 更新已有键的值
	t.Logf("after add/update, map m = %v", m)

	// 删除键值对
	delete(m, "apple") // 删除键 "apple"
	t.Logf("after delete, map m = %v", m)

	// 访问键值对
	bananaCount := m["banana"]
	t.Logf("banana count = %d", bananaCount)

	// ¥¥¥¥¥¥¥¥¥¥¥¥¥¥检查键是否存在¥¥¥¥¥¥¥¥¥¥¥¥¥
	// if 语句中同时获取值和存在性
	// 然后根据ok做判断
	if count, ok := m["apple"]; ok {
		t.Logf("apple exists with count = %d", count)
	} else {
		t.Log("apple does not exist")
	}

	// 遍历 map
	for k, v := range m {
		t.Logf("key: %s, value: %d", k, v)
	}

}
