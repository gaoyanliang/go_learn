package operator

import "testing"

// -------- 数组比较
// -------- 长度不同的两个数组不能比较，编译错误
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	//t.Log(a == c) // invalid operation: a == c (mismatched types [4]int and [5]int)
	t.Log(a == d)

}

// -------- 操作运算符和其他语言类似， + - * / %
// -------- 注意： go 中 ++ 和 -- 只能是后置的 a ++， a --。  不支持前置

// -------- 逻辑运算符和其他语言类似 && ||  ！

// -------- 位运算符和其他语言类似 & |  ^ <<  >>
// -------- 注意： go 为提高生产力，新增来一种位运算符
// -------- &^ 按位置零
// -------- 1 &^ 0 -- 1
// -------- 1 &^ 1 -- 0
// -------- 0 &^ 1 -- 0
// -------- 0 &^ 0 -- 0

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	// disable Readable Executable
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
