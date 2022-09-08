package csp

import (
	"fmt"
	"testing"
	"time"
)

/* 定义一个客户端的服务 */
func service() string {
	time.Sleep(time.Millisecond * 50)
	return "service Done"
}

/* 定义一个其它类型的服务 */
func otherTask() {
	fmt.Println("otherTask working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask is done.")
}

/* 此时调用服务都是串行的 */
func TestSequenceService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}


// --------------------

/* 创建一个具有等待机制的，阻塞性的异步服务 */
func AsyncService() chan string {
	/* 定义一个channel 里面存放的值是string类型 */
	retCh := make(chan string)

	/* 创建一个新协程，避免阻塞当前的协程 */
	go func() {
		/* 将service的调用赋值给ret */
		ret := service()
		fmt.Println("AsyncService returned result")

		/*
			将执行完成的结果放进channel
			如果service没有执行完成，则channel一直等待到执行完毕
			由于channel没有容量，所以channel的释放，必须得有人过来拿
			代码才会往下执行
		*/
		retCh <- ret
		fmt.Println("service process exited")
	}()

	/* 最后将执行完成的结果返回 */
	return retCh
}

func TestAsyncService(t *testing.T) {
	/*
		接收同步服务调用的返回值
		这时可能会出现，没有执行完成的情况
	*/
	retCh := AsyncService()

	/* 然后调用otherService */
	otherTask()

	/* 将channel的数据取出来，并打印 */
	fmt.Println(<-retCh)

}


// --------------------------

/* 创建一个具有buffer，等待机制的异步服务 */
func AsyncServiceBuffer() chan string {
	/* 给channel增加一个容量 */
	retCh := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("AsyncService returned result")

		/*
			由于channel增加了容量，所以将ret放入channel之后，代码会继续往下执行
		*/
		retCh <- ret
		fmt.Println("service process exited")
	}()
	return retCh
}

func TestAsyncServiceBuffer(t *testing.T) {
	retCh := AsyncServiceBuffer()
	otherTask()
	fmt.Println(<-retCh)
}
