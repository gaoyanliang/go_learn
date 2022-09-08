package function

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)


func TestReturnMultiValue(t *testing.T) {
	// _ 表示忽略
	a, _ := returnMultiValue()
	t.Logf("return value is %d", a)
}

/**
返回两个 10 以内的随机数
 */
func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(10);
}


// ------------- 定义函数：计算方法的耗时

func timeSpent(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now();
		ret := inner(op)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func slowFn(n int) int {
	time.Sleep(time.Second * 2)
	return n;
}

func TestTimeSpent(t *testing.T) {
	// 利用 timeSpent 包装，返回一个新的函数
	newFn := timeSpent(slowFn)

	newFn(10)
}


// ------------------ 可变参数

func sum(ops ...int) int {
	sum := 0

	for _, op := range ops {
		sum += op
	}

	return sum
}

func TestSum(t *testing.T) {
	t.Log(sum(1, 2, 3, 4, 5, 6))
	t.Log(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
}


// ------------------- 延迟函数 defer

func clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	defer clear()
	t.Log("Started")

	// panic("err")
}


