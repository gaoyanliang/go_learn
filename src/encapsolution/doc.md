## 封装

封装数据和行为

1. 结构体定义

```go
type Employee struct {
	Id string
	Name string
	Age int
}
```

2. 实例创建与初始化

```go
e := Employee{"0", "Bob", 20}
e1 := Employee{Name: "Mike", Age: 30}
e2 := new(Employee) // 注意这里返回的是引用/指针，相当于 e:= &Employee{}
e2.Id = "2" // 与其他语言的差异：通过实例的指针访问成员不需要使用 ->
e2.Name = "Rose"
e2.Age = 22
```

3. 行为（方法）定义

```go
type Employee struct {
	Id string
	Name string
	Age int
}

/* 第一种定义方式: 在实例对应方法被调用时，实例的成员会进行值复制 */
func (e Employee) String() string {
    return fmt.Sprintf("ID: %s - Name: %s - Age: %d", e.ID, e.Name, e.Age)
}

/* 通常情况下为了避免内存拷贝我们使用第二种定义方式 */
// 推荐使用
func (e *Employee) String() string {
return fmt.Sprintf("ID: %s - Name: %s - Age: %d", e.ID, e.Name, e.Age)
}
```

[Show me the code](encap_test.go)
