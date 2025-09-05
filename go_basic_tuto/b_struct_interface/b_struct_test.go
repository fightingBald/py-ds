package b_struct_interface

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

/*
组合优于继承
Go 没有继承（不搞血统学），靠组合实现复用：
struct 可以嵌套别的 struct
interface 可以嵌套别的 interface
这样避免了 OOP 常见的“继承地狱”，不需要为复用关系硬套父子类。
*/
// 1. 定义一个结构体
//要求 1: struct 字段
//定义一个 Account struct，字段设计：
//id：账号 id（public，因为需要暴露给包外使用）
//owner：账户持有人姓名（private）
//balance：余额（private，单位分，int64）
//createdAt：开户日期（只读，包外能看，不能改）
// ====== 1) 结构体：字段可见性与只读 ======
type Account struct {
	id        string    // public：对外暴露
	owner     string    // private：对外不可直接改
	balance   int64     // private：只能通过方法改
	createdAt time.Time // private：只读，getter 暴露
}

// 要求 2: 构造器
// 写一个 NewAccount(owner string, initBalance int64) (*Account, error)
// 要求检查 initBalance >= 0
// 自动生成 id（比如 "ACC-001"）
// 初始化 createdAt = 当前时间

// ====== id 生成（示例用原子计数器） ======
var accSeq int64

func nextAccountID() string {
	n := atomic.AddInt64(&accSeq, 1)
	return fmt.Sprintf("ACC-%06d", n)
}
func NewAccount(owner string, initBalance int64) (*Account, error) {
	if owner == "" {
		return nil, fmt.Errorf("owner required")
	}
	if initBalance < 0 {
		return nil, fmt.Errorf("initial balance cannot be negative")
	}
	return &Account{
		id:        nextAccountID(),
		owner:     owner,
		balance:   initBalance,
		createdAt: time.Now(),
	}, nil
}

// 要求 3: Getter / Setter
// 提供 Owner() string Getter
// 提供 SetOwner(newName string) Setter
// 提供 Balance() int64 Getter（只读 balance）
// 不提供 SetBalance（balance 只能通过 Deposit/Withdraw 改）
// 提供 CreatedAt() time.Time Getter（防御性返回拷贝）
// 读：值接收者即可（不改状态）
func (a *Account) Owner() string        { return a.owner }
func (a *Account) Balance() int64       { return a.balance }
func (a *Account) CreatedAt() time.Time { return a.createdAt }

// 写：指针接收者（会改状态）
func (a *Account) SetOwner(newName string) error {
	if newName == "" {
		return fmt.Errorf("owner cannot be empty")
	}
	a.owner = newName
	return nil
}

// 要求 4: 方法
// Deposit(amount int64) error
// amount 必须 >0
// 成功则加到 balance 上
//
// Withdraw(amount int64) error
// amount 必须 >0 且 <= balance
// 成功则扣除

// ====== 4) 行为方法（业务校验） ======
func (a *Account) Deposit(amount int64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

func (a *Account) Withdraw(amount int64) error {
	if amount <= 0 {
		return fmt.Errorf("withdraw amount must be positive")
	}
	if amount > a.balance {
		return fmt.Errorf("insufficient funds")
	}
	a.balance -= amount
	return nil
}

// 给 Account 添个 id() 方法以满足接口
func (a *Account) ID() string { return a.id }

//
//要求 6: 组合
//定义一个 Logger struct，提供方法：
//func (Logger) Log(msg string)
//再定义一个 SecureAccount：
//内嵌 Account + Logger
//重写 Withdraw 方法：先 Log 一条“安全检查”，再调用原始 Withdraw

func TestAccount(t *testing.T) {

	acc, err := NewAccount("Alice", 10000)

	if err != nil {
		t.Fatal("NewAccount error:", err)
	}

	err = acc.Deposit(5000)
	if err != nil {
		fmt.Println("Deposit error:", err)
	} else {
		fmt.Println("After deposit, balance =", acc.Balance())
	}

	err = acc.Withdraw(2000)
	if err != nil {
		fmt.Println("Withdraw error:", err)
	} else {
		fmt.Println("After withdraw, balance =", acc.Balance())
	}

	err = acc.SetOwner("Bob")
	if err != nil {
		fmt.Println("SetOwner error:", err)
	} else {
		fmt.Println("After SetOwner, owner =", acc.Owner())
	}
}
