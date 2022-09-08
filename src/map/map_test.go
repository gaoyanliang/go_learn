package _map

import "testing"

func TestMapInit(t *testing.T){
	m1 := map[int]int{1:1, 2:2, 3:3}
	t.Log(m1[2])
	t.Logf("m1 len = %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Log(m2[4])
	t.Logf("m2 len = %d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("m3 len = %d", len(m3))
}

func TestMapAccessNotExcitingKey(t *testing.T) {
	m1 := map[int]int{}
	/*不能用数组是否为空来判断，go语言会将map初始化为0 */
	t.Log(m1[1]) // 0
	m1[2] = 0
	t.Log(m1[2])

	/* 使用返回的多个来判断,map是否是空的*/
	if v, isExciting:= m1[3]; isExciting {
		t.Log(isExciting)
		t.Logf("Key 3`s value is %d", v)
	} else {
		t.Log(isExciting)
		t.Logf("Key 3 is not exciting")
	}

	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3`s value is %d", v)
	} else {
		t.Logf("Key 3 is not exciting")
	}
}

func TestTravelMsp(t *testing.T) {
	m1 := map[int]int{1:1, 2:2, 3:3}
	for k, v := range m1{
		t.Log(k, v)
	}
}
