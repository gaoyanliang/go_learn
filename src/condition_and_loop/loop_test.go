package condition_and_loop

import "testing"

// 有限循环
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n ++
	}
}

// 无限循环
func TestWhileLoop2(t *testing.T) {
	n := 5
	for {
		t.Log(n)
		n++
		//if n > 10 {
		//	break
		//}
	}
}
