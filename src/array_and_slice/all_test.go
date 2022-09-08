package array_and_slice

import (
	"fmt"
	"testing"
)

// https://golang.design/go-questions/slice/vs-array/
func TestCaseOne(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s1 从 slice 索引2（闭区间）到索引5（开区间，元素真正取到索引4），长度为3，容量默认到数组结尾，为8。
	s1 := slice[2:5]
	// s2 从 s1 的索引2（闭区间）到索引6（开区间，元素真正取到索引5），容量到索引7（开区间，真正到索引6），为5。
	s2 := s1[2:6:7]

	fmt.Printf("s1's len = %d, s1's cap = %d", len(s1), cap(s1))
	fmt.Println(s1)

	fmt.Printf("s2's len = %d, s2's cap = %d", len(s2), cap(s2))
	fmt.Println(s2)

	// 这时元素个数在容量内，s2 的修改会同时影响到 s1
	// 打印 s1 的时候，只会打印出 s1 长度以内的元素。所以，只会打印出3个元素，虽然它的底层数组不止3个元素。
	s2 = append(s2, 100)
	fmt.Printf("s1's len = %d, s1's cap = %d", len(s1), cap(s1))
	fmt.Println(s1)
	fmt.Printf("s2's len = %d, s2's cap = %d", len(s2), cap(s2))
	fmt.Println(s2)

	// 这时元素个数大于容量，触发扩容（容量翻倍）。s2 另起炉灶，将原来的元素复制新的位置，扩大自己的容量。
	// 这时 s2 的修改不会影响到 s1
	s2 = append(s2, 200)
	fmt.Printf("s2's len = %d, s2's cap = %d", len(s2), cap(s2))
	fmt.Println(s2)

	s1[2] = 20
	fmt.Printf("s1's len = %d, s1's cap = %d", len(s1), cap(s1))
	fmt.Println(s1)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)

	// output：
	//[2 3 20]
	//[4 5 6 7 100 200]
	//[0 1 2 3 20 5 6 7 100 9]
}

func TestCaseTwo(t *testing.T)  {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)

	// 5 7 9
	// 5 7 9 11
	// 5 7 9 12
}

func TestCaseThree(t *testing.T)  {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(x, 12)
	fmt.Println(s, x, y)

	// 5 7 9
	// 5 7 9 11
	// 5 7 9 11 12
}