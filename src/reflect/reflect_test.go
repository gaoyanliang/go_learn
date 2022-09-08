/*
Go 反射

reflect.TypeOf vs reflect.ValueOf

- reflect.TypeOf 返回值为 reflect.Type
- reflect.ValueOf 返回值为 reflect.Value
- 可以从 reflect.Value 获得类型
- 通过 Kind 来判断类型

特点：

- 提高了程序的灵活性
- 降低了代码的可读性
- 提升了调试难度
- 降低了程序的性能
 */

package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

/* 定义一个雇员结构 */
type Employee struct {
	EmployeeId string
	Name  string `format:"normal"`
	Age  int
}

/* 更新 Employee 年龄 */
func (e *Employee) UpdateAge(newVal int){
	e.Age = newVal
}

/* 反射获取属性 & 反射调用方法 */
func TestInvokeByName(t *testing.T){
	// 定义雇员实例
	e := &Employee{"1","Mike",30}

	// 按名字获取成员
	t.Logf("Name: value(%[1]v),Type(%[1]T)",reflect.ValueOf(*e).FieldByName("Name"))

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'name' field.")
	}else{
		t.Log("Tag:format",nameField.Tag.Get("format"))
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(6)})
	t.Log("Update Age",e)
}

func TestTypeAndValue(t *testing.T){
	var f int64 = 10
	t.Log("reflect.TypeOf", reflect.TypeOf(f))
	t.Log("reflect.TypeOf.King", reflect.TypeOf(f).Kind())
	t.Log("reflect.ValueOf", reflect.ValueOf(f))
	t.Log("reflect.ValueOf.Type", reflect.ValueOf(f).Type())
}

/* 校验类型 */
func TestBasicType(t *testing.T){
	var f float64 =12

	CheckType(f)
	CheckType(&f)
}

func CheckType(v interface{}){
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32,reflect.Float64:
		fmt.Println("Float")
	case reflect.Int,reflect.Int32,reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown",t)
	}
}


// ----------------------- 使用 reflect.DeepEqual 比较 map 和 切片 slice
func TestDeepEqual(t *testing.T) {
	a := map[int]string{1:"one",2:"two",3:"three"}
	b := map[int]string{1:"one",2:"two",4:"three"}

	//t.Log(a == b) // 函数只能与nil 比较
	t.Log(reflect.DeepEqual(a,b))

	s1 := []int{1,2,3}
	s2 := []int{1,2,3}
	s3 := []int{2,3,1}

	t.Log("s1 == s2?",reflect.DeepEqual(s1,s2))
	t.Log("s1 == s3?",reflect.DeepEqual(s1,s3))
}

// ----------------------- 万能程序
// 定义 Customer 结构。注意： Customer 和 Employee 中都存在 Name string
// 可以通过 reflect 来实现一个通用的填充属性的方法
type Customer struct {
	CookieID string
	Name string
	Age int
}

/* 填充属性方法 */
func fillBySetting(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		// Elem() 获取指针指向的值
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}

	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok bool
	)

	for k,v :=range settings {
		if field,ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		// 属性的类型需要和 value 的类型保持一致
		if field.Type == reflect.TypeOf(v){
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}

	}
	return  nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name":"Mike","Age":20}

	e := Employee{}
	if err := fillBySetting(&e,settings);err != nil{
		t.Fatal(err)
	}
	t.Log(e)

	c := new(Customer)
	if err :=fillBySetting(c,settings); err != nil{
		t.Fatal()
	}
	t.Log(*c)
}

