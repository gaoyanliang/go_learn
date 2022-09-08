## Go 函数

> 函数： 一等公民

1. 函数可以有多个返回值
2. 所有参数都是值传递： slice，map，channel 会有传引用的错觉（slice 使用的是共享空间）
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

[Show me the code](function_test.go)


### 可变参数

```go
func sum(ops ...int) int {
	sum := 0
	for _, op := range ops {
        sum += op
	}
	return sum
}
```

### 延迟执行函数 (defer 函数)

```go
func TestDefer(t *testing.T) {
	// 匿名函数 
	defer func() {
	    t.Log("Clear resources")	
    }()
	t.Log("Started")
	panic("Fatal error") // defer 仍会执行
}
```

[Show me the code](function_test.go)
