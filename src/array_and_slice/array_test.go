package array_and_slice

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4, 5}
	arr[2] = 3
	t.Log(arr, arr1, arr2)
	t.Log(arr[2], arr1[1], arr2[0])
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++{
		t.Log(arr3[i])
	}

	for _, val := range arr3{
		t.Log(val)
	}
}

func TestTwoArrayTravel(t *testing.T) {
	arr4 := [2][2]int{{1, 2}, {3, 4}}
	for _, valOut := range arr4{
		for _, val1In := range valOut {
			t.Log(val1In)
		}
	}
}

func TestArraySection(t *testing.T) {
	arr5 := [...]int{1, 2, 3, 4, 5}
	arr6 := arr5[:]
	t.Log(arr6)
}
