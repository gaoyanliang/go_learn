package main

import (
	"fmt"
	"os"
)

// 在当前目录下执行 go run hello_world.go yanliang
// output: Hello World yanliang

// 注意这里 package 和 方法名 都是 main
func main() {

	if len(os.Args) > 1 {
		fmt.Println("Hello World ", os.Args[1])
	}

	os.Exit(-1)
}
