package constant

import "testing"


/**
如何编写测试程序：
1. 源码文件以_test 结尾：xxx_test.go
2. 测试方法名以 Test 开头：func TestXXX (*testing.T) {...}
 */

/**
与其他主要编程语言的差异
1. 赋值可以进行自动类型推断
2. 在一个赋值语句中可以对多个变量进行同时赋值
 */

func TestConstan(t *testing.T) {
	t.Log("My first try!")
}

// ------------- 学习声明变量
// 实现 Fibonacci 数列  1 1 2 3 5 8 ...
func TestFib(t *testing.T) {
	// var a int = 1
	// var b int = 1

	// var {
	//   a int = 1
	//   b int = 1
	// }

	a := 1
	b := 1

	t.Log(a)
	for i := 0; i < 10; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log(" end")
}

// ------------- 学习如何定义变量

const (
	Mon = 1 + iota
	Tue
	Wed
	Tus
	Fir
)

func TestCon1(t *testing.T) {
	t.Log(Mon, Tue, Wed, Tus, Fir)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCon2(t *testing.T) {
	a := 7 // 0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
