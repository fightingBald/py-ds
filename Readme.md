# py-ds

项目用于存放 Go 语言的基础练习与小工具，主要用于学习数据结构、基础语法和单元测试实践。

## 项目结构

- go.mod — Go 模块描述
- Readme.md — 项目说明（本文件）
- counter/
  - counter.go — 计数器示例实现
  - counter_test.go — 计数器单元测试
- go_basic_tuto/
  - a_basic_type/ — 基础类型与集合示例
    - a_list.go
    - a_map_test.go

> 注：仓库结构尽量保持简单，便于按主题分组学习与测试。

## 快速开始

1. 克隆仓库并进入项目根目录。
2. 运行测试：

   go test ./...


0. 预热（0.5 天）

装：Go 1.22+（你已看 1.24 可用）、go mod, go test, go vet, golangci-lint

读：Effective Go（扫一遍风格），Tour of Go（把语法按钮全点完）

认识：包管理只用 go mod，无需虚拟环境；单一可执行文件，跨平台无痛。

Python 对照：pip + venv → go mod；pytest → go test；black/flake8 → gofmt/golangci-lint。

1. 语言基石（3–5 天）

目标：写出“干净的 Go 函数”，搞懂值/指针、切片/映射、接口、错误。

必须掌握

值与指针：何时用 *T（大对象/需原地修改/同步原语），何时用值拷贝。

slice/map：容量 cap、扩容、nil map 不能写、range 是复制元素的“快照”语义（1.22 修复 for 变量捕获坑）。

接口：隐式实现，小接口（如 io.Reader）优于巨物。

错误：返回值而非异常；errors.Is/As/Join；哨兵错误 vs 包装。

并发：goroutine、channel、context.Context 取消、sync.WaitGroup、atomic。

Python 对照速记表（超精简）

主题	Python	Go
包管理	pip/venv	go mod init/tidy
异常/错误	try/except	val, err := f() + if err != nil
OOP	动态鸭子类型	结构体 + 接口（隐式）
并发	threading/asyncio	goroutine + channel + context
列表/字典	list/dict	slice/map（注意零值与容量）
日志	logging	log/slog（Go 1.21+ 标准）
测试	pytest	go test -v ./...；_test.go 文件

小练习（30–60 分钟）

写一个并发网页探活器：输入一堆 URL，最多 20 并发，context.WithTimeout 5s，统计成功率 & 平均耗时。
升级：加入 -workers=N、-json 输出。
你会踩到：http.Client 复用、defer resp.Body.Close()、channel 泄漏。