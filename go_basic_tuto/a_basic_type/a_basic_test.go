package a_basic_type

import "testing"

func TestTheVariabletype(t *testing.T) {
	const a int = 10
	var b int = 20
	c := 30
	t.Log(a, b, c) // 10 20 30
	t.Logf("Type of a is %T, value is %v", a, a)
}

func TestIfStateMention(t *testing.T) {
	a := 10
	if a > 5 {
		t.Log("a is greater than 5")
	} else {
		t.Log("a is not greater than 5")
	}

	if b := a * 2; b > 15 { // if 里可以初始化变量
		t.Log("b is greater than 15")
	} else {
		t.Log("b is not greater than 15")
	}
}

func TestLoopStatement(t *testing.T) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	t.Log("Sum of first 10 numbers is", sum)

	for sum < 100 {
		sum += 10
	}

	arr := []string{"apple", "banana", "cherry"}
	for index, value := range arr {
		t.Logf("Index: %d, Value: %s", index, value)
	}

}

func swap(x, y int) (int, int) {
	return y, x
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	t.Logf("Before swap: a=%d, b=%d", a, b)
	a, b = swap(a, b)
	t.Logf("After swap: a=%d, b=%d", a, b)
}
