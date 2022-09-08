package _type

import (
	"math"
	"testing"
)

// -- 定义别名
type MyInt int64

// ------------ go 语言不允许隐式类型转换
// func TestType(t *testing.T) {
// 	var a int = 1
// 	var b int64
//
// 	b = a
// 	t.Log(a, b)
//
// 	a = b
// 	t.Log(a, b)
// }

// ------------ 别名和原有类型也不能进行隐式类型转换
// func TestType2(t *testing.T) {
// 	var a int = 1
// 	var b int64
//
// 	b = int64(a)
//
// 	var c MyInt
// 	c = b
// 	t.Log(c)
//
// }

// ----------- 类型转换只能显式转换
func TestType3(t *testing.T) {
	var a int32 = 1
	var b int64

	b = int64(a)

	t.Log(a, b)

}

// ----------- 类型的预定义值
func TestType4(t *testing.T) {
	t.Log(math.MinInt)
	t.Log(math.MaxInt)
}

// ---------- go 支持指针，但是不能通过指针来进行运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

// ----------- 打印初始值

func TestInitValue(t *testing.T) {
	// string 类型 初始值为 nil
	var s string
	t.Log("*" + s + "*")
	t.Log("string init value is", len(s))

	var a int
	t.Log(a)

	var b int64
	t.Log(b)

	var c byte
	t.Log(c)

}
