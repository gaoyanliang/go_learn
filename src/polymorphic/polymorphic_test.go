/* 多态 */
package polymorphic

import (
	"fmt"
	"testing"
)

/* 定义一个自定类型 （别名） */
type Code string

/* 定义一个接口 */
type Programmer interface {
	WriteHelloWorld() Code
}

/* 定义一个类型 */
type GoProgrammer struct {
}

/* 给类型绑定一个方法 */
func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello,World\")"
}

/* 定义一个类型 */
type JavaProgrammer struct {
}

/* 给类型绑定一个方法 */
func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello,World\")"
}

/* 创建一个方法来打印，对应的类型，以及类型绑定的方法 */
func writeFirstProgram(p Programmer) {
	/* 传入什么类型，打印相应类型绑定的方法 */
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphic(t *testing.T) {
	/* 创建接口类型实例必须是指针类型 */
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}