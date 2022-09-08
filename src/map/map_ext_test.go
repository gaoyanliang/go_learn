package _map

import "testing"

// map 的 value 可以是一个方法
func TestMapWithFuncValue(t *testing.T)  {
	m := map[int]func(op int) int{}

	m[0] = func(op int) int {return op}
	m[1] = func(op int) int {return op * op}
	m[2] = func(op int) int {return op * op * op}

	t.Log(m[0](2), m[1](2), m[2](2))
}

// 使用 Map 实现 Set
func TestMapFoSet(t *testing.T) {
	// 1. 定义一个 map
	m := map[int]bool{}
	m[0] = true;
	m[1] = true;

	n := 0
	if m[n] {
		t.Logf("%d is exsiting.", n)
	} else {
		t.Logf("%d is not exsiting.", n)
	}
	t.Logf("m's len is %d", len(m))
	t.Log("delete m[0]")
	delete(m, 0)

	if m[n] {
		t.Logf("%d is exsiting.", n)
	} else {
		t.Logf("%d is not exsiting.", n)
	}
}