/*
接口定义的原因，两个模块可能需要使用相同的功能
但是具体实现不一样。此操作是为了方便扩展。
且接口的实现不依赖接口的定义，采用duck type方式
*/
package _interface

import (
	"fmt"
	"testing"
	"time"
)

/* 首先需要定义一个接口 interface */
type Program interface {
	WriteHelloWorld() string
}

// 定义类型
type GoProgrammer struct {}

// 类型作为参数，绑定方法/实例
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello,World\")"
}

func TestClient(t *testing.T) {
	/* 创建一个接口类型变量 */
	var g Program
	/* 使用接口需要传指针 */
	g = new(GoProgrammer)
	t.Log(g.WriteHelloWorld())
}


// ------------- 自定义类型

/* 由于参数比较长，所以我们创建一个别名*/
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("time spent :", time.Since(start))
		return ret
	}
}

/* 创建一个停止10s的函数 */
func slowFn(op int) int {
	time.Sleep(time.Second * 2)
	return 2
}

func TestFn(t *testing.T) {
	/* 将函数的计时功能传递给新的值 */
	newTimeSpendSlowFn := timeSpent(slowFn)
	t.Log(newTimeSpendSlowFn(2))
}