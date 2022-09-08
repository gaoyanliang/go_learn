/* 此篇为go语言数据与行为的封装 */
package encapsolution

import (
"fmt"
"testing"
"unsafe"
)

// ---------- 结构体的定义 & 对象的创建与初始化
type Employee struct {
	ID   string
	Name string
	Age  int
}

func TestCreatEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 1}
	e1 := Employee{Name: "Mike", Age: 20}
	e2 := new(Employee) // 返回的指针类型
	e2.ID = "2"
	e2.Name = "Rose"
	e2.Age = 10
	t.Log(e1.ID)
	t.Log(e, e1, e2)
	// %T 返回类型
	t.Logf("e`s type is %T ", e)
	t.Logf("&e`s type is %T ", &e)
	t.Logf("e1`s type is %T ", e1)
	t.Logf("e2`s type is %T", e2)
}

// --------------- 行为（方法）的定义

/* 第一种定义方式在实例对应方法被调用时，实例的成员会进行值复制 */
func (e Employee) String1() string {
	/* 查看传入参数的地址 */
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID is %s, Name is %s, Age is %d", e.ID, e.Name, e.Age)
}

func TestStructOperation1(t *testing.T) {
	e := &Employee{"0", "Bob", 1}
	/* 可以看见，此时内存地址不一致，因此判断发生了数据的复制 */
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	/* 在go里面，对于指针类型的变量
	也不用去 -> 去访问，可以直接使用 "." 去访问
	*/
	t.Log(e.String1())
}

/* 通常情况下为了避免内存拷贝我们使用第二种定义方式 */
func (e *Employee) String2() string {
	/* 查看传入参数的地址 */
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID is %s, Name is %s, Age is %d", e.ID, e.Name, e.Age)
}

func TestStructOperation(t *testing.T) {
	e := &Employee{Name: "Mike", Age: 20}
	/* 查看原处的地址 */
	/* 可以看见，内存地址是一样的，说明并没有数据被拷贝*/
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	/* 在go语言指针和指，都可以作为参数来访问 */
	t.Log(e.String2())
}