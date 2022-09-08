package series

import (
	"errors"
	"fmt"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

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

/* 小写的函数将不能被访问 */
func square(n int) int {
	return n * n
}

func Square(n int) int {
	return n * n
}

/*
init 会优先于main先执行
init的执行顺序根据包的包的导入顺序执行
每个包可以有多个init函数
甚至每个源文件也可以有多个init函数
*/
func init() {
	fmt.Println("init1......")
}

func init() {
	fmt.Println("init2......")
}
