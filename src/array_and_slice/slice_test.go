package array_and_slice

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	/* slice 切片
	与数组的差别在于[]不填
	slice是可变长度的
	*/
	var sli0 []int
	/* capacity 容量 */
	t.Log(len(sli0), cap(sli0))

	sli0 = append(sli0, 1)
	t.Log(len(sli0), cap(sli0))

	// length为4,表示可访问的元素为4个， capacity为4
	sli1 := []int{1, 2, 3, 4}
	t.Log(len(sli1), cap(sli1))

	// 使用make生成slice,
	//len表示长度。同时也是元素可访问的有几个、初始化的个数
	// 而cap仅仅表示容量
	sli2 := make([]int, 3, 5)
	t.Log(len(sli2), cap(sli2))
	t.Log(sli2[1], sli2[2], sli2[0])
	sli2 = append(sli2, 1)
	t.Log(len(sli2), cap(sli2))
	t.Log(sli2[1], sli2[2], sli2[0], sli2[3])
}

func TestSliceGrowing(t *testing.T) {
	var sli0 []int
	t.Log(len(sli0), cap(sli0))
	for i := 0; i < 20; i++{
		sli0 = append(sli0, i)
		t.Log(len(sli0), cap(sli0))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul",
		"Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T){
	sli0 := []int{1,2,3,4}
	//sli1 := []int{1,2,3,4}
	if sli0 == nil {
		t.Log("equal")
	}
}

func TestSlicePrint(t *testing.T) {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// s1 从 slice 索引2（闭区间）到索引5（开区间，元素真正取到索引4），长度为3，容量默认到数组结尾，为8。
	s1 := slice[2:5]

	fmt.Printf("s1: %v", s1)
	fmt.Println()
	fmt.Println(s1)

	s2 := make([]int, 3, 4)
	s2[3] = 100
	fmt.Println(s2)
}
