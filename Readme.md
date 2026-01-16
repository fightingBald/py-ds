# mp_project Go 语法小练习

这个目录里放了一个可独立运行的小练习，逻辑与语法点对应你贴的 Go 函数。
重点覆盖：

- 多返回值（比如 `(*UserProfile, error)`）
- 短变量声明（`:=`）与错误检查
- 早返回（`if err != nil { return nil, err }`）
- 方法链式调用（builder 风格）
- `context.Context` 作为第一个参数
- 字符串裁剪与布尔标记

文件说明：
- `mp_project/profile/support.go`：简化后的类型与辅助函数。
- `mp_project/profile/exercise.go`：需要你完成的 TODO。
- `mp_project/profile/exercise_test.go`：期望行为的测试用例。

使用方式：
1. 打开 `mp_project/profile/exercise.go`，实现所有 TODO。
2. 运行 `go test ./mp_project/...`。

Python 对照小抄：
- `x, err := f()` 类似 Python 的“拆包 + 显式错误检查”。
- `if err != nil { return nil, err }` 类似错误即早返回。
- 方法链（`Check(...).CanDoAny(...).OnResource(...)`）是 builder 模式。
