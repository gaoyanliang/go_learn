package anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/* 定义一个任务 */
func runTask(i int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", i)
}

/* 多个运行任务，只要有一个返回，就结束 */
func FirstRespoonse() string {
	numOfRuner := 10
	// 注意：这里要创建有 buffer 的 channel。 思考为什么？
	ch := make(chan string, numOfRuner)

	for i := 0; i < numOfRuner; i ++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestResponse(t *testing.T) {
	t.Log("Before: ", runtime.NumGoroutine())
	t.Log(FirstRespoonse())
	time.Sleep(time.Second * 1)
	t.Log("After: ", runtime.NumGoroutine())
}

