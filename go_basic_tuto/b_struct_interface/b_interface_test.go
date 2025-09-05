package b_struct_interface

import (
	"time"
)

/*
在 Go 里：只要方法对得上，struct 就自动实现接口。
👉 好处：
使用方依赖接口，不依赖具体实现。
实现方不需要提前知道接口的存在。
Go 鼓励 小接口（只定义一两个方法），像 io.Reader、io.Writer。
小接口可以灵活组合，避免臃肿的大接口（Java 里经常见到）。
👉 好处：
用户只需要实现自己需要的最小功能。
函数参数可以要求非常精确的行为。

你写了一个新 struct，只要方法名对得上，就自然满足接口。
老接口调用处直接能用，不需要额外声明“我实现了这个接口”。
👉 举个例子：
你写了个新类型 KafkaReader，只要它有 Read 方法，它就自动是 io.Reader，你都不用 import io 包。
这就是所谓的 隐式接口实现。
*/

// ====== 5) 只读接口（限权视图） ======
//✅ 限权 / 封装
//你不希望外部代码直接改 Account 的内部状态（例如直接调用 Deposit / Withdraw）。
//但又想把一些字段暴露出去让人查（id、Owner、Balance）。
//这时候就返回一个 ReadOnlyAccount 接口，调用方拿到的就是“只能看，不能改”的版本。
// func GetAccountView(id string) ReadOnlyAccount
//- 返回 *Account → 外部能读+写。
//- 返回 ReadOnlyAccount → 外部只能读。
//外部用的人 即使知道这是 Account，因为函数签名返回的是 ReadOnlyAccount，他们只能调用 getter 方法，看不到 Deposit / Withdraw。
//👉 这就是信息隐藏 / API 限权。
//✅ 解耦 / 可替换实现
//将来你可能有别的实现不一定是 Account（例如只读快照、远程代理）。
//只要满足 ReadOnlyAccount，调用方完全不用关心背后实现是什么。
//👉 这是 依赖倒置原则（Depend on abstractions, not on concretions）的体现。

type ReadOnlyAccount interface {
	ID() string // 注意：我们给 Account 实现一个 ID() 方法
	Owner() string
	Balance() int64
	CreatedAt() time.Time
}

/*
✅ Mock / 测试
你在单元测试里，可以随便写一个假的实现，只需要满足 ReadOnlyAccount 接口。
不需要完整实现 Account 的所有逻辑。
*/
type FakeReadOnlyAccount struct{}

func (f FakeReadOnlyAccount) ID() string           { return "FAKE" }
func (f FakeReadOnlyAccount) Owner() string        { return "Test" }
func (f FakeReadOnlyAccount) Balance() int64       { return 9999 }
func (f FakeReadOnlyAccount) CreatedAt() time.Time { return time.Now() }

func returnAccountView(id string) ReadOnlyAccount {
	// 省略查找逻辑，直接返回一个 Account
	acc, _ := NewAccount("Alice", 10000)
	return acc // 隐式转换为 ReadOnlyAccount
}
