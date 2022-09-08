package error

import (
	"fmt"
	"os"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	/*
		调用panic之前，会调用defer
		Exit 则不会运行defer
	*/
	defer func() {
		fmt.Println("Finally!")
	}()


	/*
		可以结合defer 和 recover来恢复
		recover一定要小心使用，有时或许真的发生了系统资源被沾满的情况
		这时不如 crash 掉，系统会帮我重启，来恢复不确定性的错误
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover from", err)
		}
	}()

	fmt.Println("Start")

	os.Exit(-1)

	/* panic接收的一个空接口，输出调用栈的信息 */
	//panic(errors.New("something wrong"))
}

