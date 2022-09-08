## 闭包

闭包：引用了外部变量的匿名函数

> 函数 + 引用外部变量 = 闭包   注意：这里的外部变量也可以是一个函数

原理：

1. 函数可以作为返回值
2. 函数内查找参数的顺序，现在自己内部找，找不到到外层找

```go
package main

import "fmt"

func main()  {
	// 定义一个字符串
	str := "www.quanxiaoha.com"

	// 创建一个匿名函数
	function := func() {
		// 给字符串 str 赋予一个新的值，注意: 匿名函数引用了外部变量，这种情况形成了闭包 
		str = "犬小哈教程"
		// 打印
		fmt.Println(str)
	}

	// 执行闭包
	function()
}

// output: 犬小哈教程
```

## 闭包的记忆效应

闭包在引用外部变量后具有记忆效应，闭包中可以修改变量，变量会随着闭包的生命周期一直存在，此时，闭包如同变量一样拥有了记忆效应。

```go
package main

import "fmt"

// 定义一个累加函数，返回类型为 func() int, 入参为整数类型，每次调用函数对该值进行累加
func Add(value int) func() int  {
	// 返回一个闭包
	return func() int {
		// 累加
		value++
		// 返回累加值
		return value
	}
}

func main()  {
	// 创建一个累加器，初始值为 1
	accumulator := Add(1)

	// 累加1并打印
	fmt.Println(accumulator())
	// 再来一次
	fmt.Println(accumulator())

	// 创建另一个累加器，初始值为 10
	accumulator2 := Add(10)

	// 累加1并打印
	fmt.Println(accumulator2())
}

// output:
// 2
// 3
// 11
```

## 通过闭包实现一个生成器

可以通过闭包的记忆效应来实现设计模式中工厂模式的生成器。下面的代码示例展示了创建游戏玩家生成器的过程。

```go
package main

import "fmt"

// 定义一个玩家生成器，它的返回类型为 func() (string, int)，输入名称，返回新的玩家数据
func genPlayer(name string) func() (string, int)  {
	// 定义玩家血量
	hp := 1000
	// 返回闭包
	return func() (string, int) {
		// 引用了外部的 hp 变量, 形成了闭包
		return name, hp
	}
}

func main()  {
	// 创建一个玩家生成器
	generator := genPlayer("犬小哈")

	// 返回新创建玩家的姓名, 血量
	name, hp := generator()

	// 打印
	fmt.Println(name, hp)
}

// 犬小哈 1000
```

> **闭包具有面向对象语言的特性 —— 封装性，变量 hp 无法从外部直接访问和修改。**