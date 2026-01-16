# 学生数据内存案例：Go 数据结构与指针循序渐进教程

这是一个 **贴近真实场景** 的练习：内存里有一批学生数据，你要实现各种常见操作。
每道题都在 `exercise.go` 中给出 **中文提示**，说明思路与关键函数/操作。

## 目标
- 熟悉 Go 里最常用的数据结构：slice、map、struct、指针
- 强制建立“面向 nil 编程”的习惯
- 覆盖常用操作与常见坑（奇技淫巧）

## 目录
- `common/`：抽离的公共结构体定义（Student/Profile/Contact）
- `exercise.go`：题目 + 中文提示（TODO）
- `exercise_test.go`：测试用例（定义期望行为）

## 使用方式
1. 打开 `mp_project/student_ds_tutorial/exercise.go`，按顺序实现 TODO。
2. 运行 `go test ./mp_project/student_ds_tutorial -v`。

## 学习路线（按顺序做）

### 1) 字符串处理（循序渐进）
题目：`NormalizeName`、`MakeStudentEmail`、`ParseTagLine`、`BuildDisplayName`、`HasNamePrefix`、`SplitFullName`、`NormalizeStudentID`、`IsSchoolEmail`、`NameContainsKeyword`、`NameToSlug`、`SplitSubjects`
关键函数/操作：
- `strings.TrimSpace` / `strings.Fields` / `strings.Join`
- `strings.ToLower` / `strings.HasPrefix` / `strings.HasSuffix` / `strings.Contains`
- `strings.TrimPrefix` / `strings.TrimSuffix` / `strings.ReplaceAll` / `strings.Split`
- `strings.Builder` 做高效拼接
奇技淫巧：
- `strings.Fields` 可以自动压缩多余空白

### 2) Slice（切片）常用操作
题目：`FilterEnrolled`、`Names`、`FindByID`、`AddStudent`、`UpdateStudentName`、`RemoveByID`、`InsertAt`、`CloneStudents`、`DedupByID`、`SortByName`、`ContainsID`、`DeleteAt`、`ReverseIDs`、`CompactIDs`
关键函数/操作：
- `len` / `append` / `copy`
- 切片删除：`append(s[:i], s[i+1:]...)` / `slices.Delete`
- 插入：`append` + `copy` 或 `slices.Insert`
- 常用工具：`slices.Clone` / `slices.Reverse` / `slices.Compact` / `slices.ContainsFunc`
奇技淫巧：
- range 变量是 **拷贝**，不能 `return &v`（见 `FindByID`）

### 3) Map（字典）常用操作
题目：`IndexByID`、`GetFromIndex`、`UpsertIndex`、`GroupByClass`、`CountByGrade`、`MergeScoreTotals`、`DeleteFromIndex`、`ScoreOf`、`IDsFromIndex`、`CloneScoreMap`、`CopyScoreMap`、`DeleteLowScores`、`ScoreSubjectsSorted`
关键函数/操作：
- `make(map[K]V)` / `m[key]` / `delete(m, key)`
- `value, ok := m[key]`（逗号 ok）
- 常用工具：`maps.Clone` / `maps.Copy` / `maps.DeleteFunc` / `maps.Keys`
奇技淫巧：
- 读 nil map 不会 panic，写 nil map 会 panic
- `delete(nilMap, key)` 安全但不会改变什么
- map 索引到指针时，避免对 range 变量取地址

### 4) 指针与 nil
题目：`NewStudent`、`EnsureProfile`、`UpdatePhone`、`AddScore`、`EmergencyName`
关键函数/操作：
- `&T{}` / `new(T)` / `nil` 判断
- 指针接收者方法的 nil 安全

### 5) 数据类型转换
题目：`ParseGrade`、`FormatGrade`、`ScoresToStringMap`、`ParseScores`、`IDsToCSV`、`CSVToIDs`
关键函数/操作：
- `strconv.Atoi` / `strconv.Itoa`
- `strings.Split` / `strings.Join`
- map 与 slice 互转

### 6) Set 操作
题目：`TagSet`、`SetUnion`、`SetIntersect`、`SetDiff`、`SetToSortedSlice`、`CommonTags`
关键函数/操作：
- set 表示：`map[string]struct{}`
- 交并差：遍历 + 存在性判断
- set 转切片排序：`sort.Strings`

### 7) 排序与搜索（自定义排序）
题目：`UniqueTags`、`UniqueClasses`、`SortByScore`、`SortByGradeThenName`、`StableSortByClass`、`BinarySearchByID`、`TopStudentBySubject`
关键函数/操作：
- 自定义排序：`slices.SortFunc` / `sort.SliceStable`
- 二分查找：`slices.BinarySearchFunc`
- 最大值搜索：循环 + 逗号 ok

## nil 编程清单（强制习惯）
- 对 `*T` 操作前先判断 `nil`
- 写 map 前先 `make`
- nil slice 可以直接 `append`
- 返回 slice 时如需避免外部改动，用 `copy`/`clone`
