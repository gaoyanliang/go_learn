# Map

- [超详细 Map 实现分析](https://qcrao.com/post/dive-into-go-map/)

## Demo


```go
// Map 声明
m := map[string]int{"one": 1, "two": 2, "three": 3}

m1 := map[string]int{}
m1["one"] = 1

m2 := make(map[string]int, 10 /* Initial Capacity */)
// 为什么不初始化 len？

// 对比 slice, len 所在的位置会初始化为零值，map 无法初始化
```


### Map 访问

在访问的 Key 不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在

### Map 遍历

```go
m := map[string]int{"one": 1, "two": 2, "three": 3}

for k,v := range m {
	t.Log(k, v)
}
```

[Show me the code]()

## Map 与 工厂模式

- Map 的 value 可以是一个方法
- 与 Go 的 Dock Type 接口方式一起，可以方便的实现单一方法对象的工厂模式

### 使用 Map 实现 Set

Go 的内置集合中没有 Set 实现，可以使用 `map[type]bool` 来实现

1. 元素的唯一性
2. 基本操作
   1. 添加元素
   2. 判断元素是否存在
   3. 删除元素
   4. 元素个数

