/* 错误处理 */
package error

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"testing"
)

/* 通常情况下，会自定义一个错误类型的变量 */
var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

/* 对程序新增错误处理 */
func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThanHundredError
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

/*
嵌套输出例子
实际编写错误处理要避免嵌套
*/
func GetFibonacciCall(str string) {
	var (
		i    int
		err  error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println(err)
	}
}

/* 有错误抛出错误信息，并退出 */
func GetFibonacciCall1(str string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}

	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}

	/* 除非完全没有错误，才打印list */
	fmt.Println(list)

}

func TestGetFibonacci(t *testing.T) {
	/* 我们可以通过调用的判断来执行，相应的打印 */
	if v, err := GetFibonacci(101); err != nil {
		if err == LessThanTwoError {
			fmt.Println("n is less")
		} else if err == LargerThanHundredError {
			fmt.Println("n is larger")
		}
	} else {
		t.Log(v)
	}

	/* 相应的调用时，也采用多返回值的方式 */
	/*if v, err := GetFibonacci(-10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
	if v, err := GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}*/

	//t.Log(GetFibonacci(-10))
}

func TestGetFibonacciCall(t *testing.T) {
	/* 实际两者功能都是一样的，只是第二个有错误就抛出，简化了思考成本 */
	GetFibonacciCall("101")
	GetFibonacciCall1("10")
}


