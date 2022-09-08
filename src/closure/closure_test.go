package closure

import (
	"fmt"
	"testing"
)

func f1(x int) func(int) int {
	fmt.Println("f1 x = ", x)
	return func(y int) int {
		x += y
		fmt.Println(x)
		return x
	}
}

func TestClosure1(t *testing.T) {
	var x int = 100
	closure := f1(x)

	closure(200)
	fmt.Println("x = ", x)
	closure(300)
	fmt.Println("x = ", x)
}

// -----------------
func f2(f func()) {
	fmt.Println("this is f2")
	f()
}

func f3(x, y int) int {
	fmt.Println("this is f3, cal x + y = ", x + y)
	return x + y
}

// 要求实现 f2(f3)
// 由于f3 不符合 f2 的参数类型，可以通过闭包（使用其他函数进行包装）实现

func f4(f func(x, y int) int, x, y int) func() {
	tmpf := func() {
		f(x, y)
	}
	return tmpf
}

func TestClosure2(t *testing.T)  {
	f2(f4(f3, 100, 200))
}

// -------------------------
// 闭包案例
func f5(base int) (func(x int) int, func(x int) int) {
	add := func(x int) int {
		base += x
		return base
	}

	sub := func(x int) int {
		base -= x
		return base
	}
	return add, sub
}

func TestClosure3(t *testing.T)  {
	add , sub := f5(10)
	fmt.Println(add(1), sub(2)) // 11  9
	fmt.Println(add(3), sub(4)) // 12  8
	fmt.Println(add(5), sub(6)) // 13  7
}
