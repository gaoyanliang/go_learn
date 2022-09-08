// Go 不支持继承
package extension

import (
	"fmt"
	"testing"
)

// 定义一个 Pet（以下面向对象来说，都只是绑定关系）
type Pet struct{}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("", host)
}

// ------------------------------
/*
以下为go的扩展功能
以我自己的理解，就是往外包了一层
*/
type Dog struct {
	p *Pet
}

/* 我们也可以实现重载 */
func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

/*
但是如果想要通过调用speak_to，来调用修改后的speak
这时我们就需要我们自己来定义speak_to
*/
func (d *Dog) SpeakTo(host string) {
	d.p.SpeakTo(host)
}

func TestDog(t *testing.T) {
	d := new(Dog)
	d.Speak()
	d.SpeakTo("Chao")
}

// ---------------------------
/* 通过匿名嵌套类型 */
type Duck struct {
	Pet
}

func (d *Duck) Speak() {
	fmt.Println("Wang!")
}

/* 来实现所谓的继承关系 */
func TestDuck(t *testing.T) {
	//var duck Pet = new(Duck) // Pet类型不能声明为Duck类型
	d := new(Duck)
	/*  */
	d.Speak()
	d.SpeakTo("Chao")
}
