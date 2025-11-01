package go_practice

import (
	"slices"
	"strconv"
	"strings"
)

// DSPlayground 用来模拟各种数据结构操作的练习题场景。
// Numbers 代表一组需要增删改查、过滤、映射的整数；
// Users 代表一个基本的用户表，方便练习查找、分组和更新；
// Tags 用 map 当作 set，帮助熟悉集合并、差、交操作。
type DSPlayground struct {
	Numbers []int
	Users   []UserProfile
	Tags    map[string]struct{}
}

// UserProfile 描述一个用户的最小信息集合，用于做 CRUD 和过滤练习。
type UserProfile struct {
	ID    int
	Name  string
	Score int
}

// 题目：完成 NewDSPlayground，把 tags 切片转换成去重的 set，并返回练习场景对象。
func NewDSPlayground(nums []int, users []UserProfile, tags []string) *DSPlayground {
	numsCopy := append([]int(nil), nums...)
	usersCopy := append([]UserProfile(nil), users...)

	setTags := map[string]struct{}{}

	for _, tag := range tags {
		setTags[tag] = struct{}{}
	}

	return &DSPlayground{Numbers: numsCopy, Users: usersCopy, Tags: setTags}

}

// 题目：实现 AppendUniqueNumber，把新数字追加到 Numbers 中，要求已存在则不追加。
func AppendUniqueNumber(ds *DSPlayground, n int) {
	for _, num := range ds.Numbers {
		if num == n {
			return
		}
	}
	ds.Numbers = append(ds.Numbers, n)
}

// 题目：实现 RemoveNumberAt，删除 Numbers 中指定下标的元素，越界时保持原样。
func RemoveNumberAt(ds *DSPlayground, index int) {
	if index < 0 || index >= len(ds.Numbers) {
		return
	}
	ds.Numbers = append(ds.Numbers[:index], ds.Numbers[index+1:]...)

}

// 题目：实现 FilterNumbers，把小于 min 或大于 max 的数过滤掉，保留闭区间内的数。
func FilterNumbers(ds *DSPlayground, min, max int) []int {
	numsFiltered := make([]int, 0, len(ds.Numbers))

	for _, num := range ds.Numbers {
		if num >= min && num <= max {
			numsFiltered = append(numsFiltered, num)
		}
	}

	return numsFiltered
}

// 题目：实现 MapNumbersToString，把 Numbers 转换成字符串切片，例如 []int{1,2} -> []string{"#1","#2"}。
func MapNumbersToString(ds DSPlayground) []string {
	result := make([]string, len(ds.Numbers))
	for i, num := range ds.Numbers {
		numStr := strconv.Itoa(num)
		result[i] = "#" + numStr
	}
	return result
}

// 题目：实现 UpsertUserScore，根据 ID 更新用户分数；若用户不存在则追加新用户。
func UpsertUserScore(ds *DSPlayground, id int, name string, score int) {
	for i := range ds.Users {
		if ds.Users[i].ID == id {
			ds.Users[i].Name = name
			ds.Users[i].Score = score
			return
		}
	}
	ds.Users = append(ds.Users, UserProfile{ID: id, Name: name, Score: score})
}

// 题目：实现 SortUsersByScoreThenID，就地排序 Users，按 Score 从高到低，同分按 ID 从小到大。
func SortUsersByScoreThenID(ds *DSPlayground) {
	if ds == nil || len(ds.Users) < 2 {
		return
	}
	slices.SortFunc(ds.Users, func(a, b UserProfile) int {
		switch {
		case a.Score > b.Score:
			return -1
		case a.Score < b.Score:
			return 1
		case a.ID < b.ID:
			return -1
		case a.ID > b.ID:
			return 1
		default:
			return 0
		}
	})
}

// 题目：实现 DeleteUserByNamePrefix，删除 Name 以 prefix 开头的用户，保持顺序稳定。
func DeleteUserByNamePrefix(ds *DSPlayground, prefix string) {
	if prefix == "" {
		return
	}
	filtered := ds.Users[:0]
	for _, user := range ds.Users {
		if strings.HasPrefix(user.Name, prefix) {
			continue
		}
		filtered = append(filtered, user)
	}
	ds.Users = filtered
}

// 题目：实现 FilterUsersByScore，返回所有分数在区间 [min,max] 的用户副本。
func FilterUsersByScore(ds DSPlayground, min, max int) []UserProfile {
	result := make([]UserProfile, 0, len(ds.Users))

	for _, user := range ds.Users {
		if user.Score < min || user.Score > max {
			continue
		}
		result = append(result, user)
	}
	return result
}

// 题目：实现 CollectUserNames，把用户姓名按字典序排序后返回，重复姓名只保留一次。
// 直接拉一遍名字，先排序再用 Compact 去掉相邻重复，省掉 map。
func CollectUserNames(ds *DSPlayground) []string {
	if len(ds.Users) == 0 {
		return nil
	}
	names := make([]string, len(ds.Users))
	for i, user := range ds.Users {
		names[i] = user.Name
	}

	slices.Sort(names)
	names = slices.Compact(names)
	return names

}

// 题目：实现 MergeTags，把新的标签切片合并进 Tags set，并返回新增数量。
func MergeTags(ds *DSPlayground, newTags []string) int {
	originalTagsLen := len(ds.Tags)

	for _, tag := range newTags {
		ds.Tags[tag] = struct{}{}
	}
	return len(ds.Tags) - originalTagsLen
}

// 题目：实现 IndexUsersByScore，把分数映射到用户 ID 列表，例如 100 -> []int{1,3}。
func IndexUsersByScore(ds DSPlayground) map[int][]int {
	userScoreId := make(map[int][]int, len(ds.Users))
	for _, user := range ds.Users {
		userScoreId[user.Score] = append(userScoreId[user.Score], user.ID)
	}
	return userScoreId

}

// 题目：实现 SplitTags，把 Tags 分成两个切片：长度为偶数的标签和长度为奇数的标签。
func SplitTags(ds DSPlayground) (even []string, odd []string) {
	even = make([]string, 0, len(ds.Tags))
	odd = make([]string, 0, len(ds.Tags))

	for tag, _ := range ds.Tags {
		if len(tag)%2 == 0 {
			even = append(even, tag)
		} else {
			odd = append(odd, tag)
		}
	}
	return even, odd
}
