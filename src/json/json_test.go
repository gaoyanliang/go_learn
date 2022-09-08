/*
内置的 JSON 解析
利用反射实现，通过 FiledTag 来标识对应的 json 值

性能较低，生产环境不推荐使用。

type BasicInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo JobInfo `json:"job_info"`
}

更快的 JSON 解析

EasyJSON 采用代码生成 而非反射

安装： go get -u github.com/mailru/easyjson
使用： easyjson -all <结构定义>.go

 */
package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 定义一个 json 字符串
var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr),e)

	if err != nil{
		t.Error(err)
	}

	fmt.Println(e)

	if v,err := json.Marshal(e);err == nil {
		fmt.Println(string(v))
	}else{
		t.Error(err)
	}
}

func TestEasyJson(t *testing.T) {
	e := Employee{}
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)

	if v,err := e.MarshalJSON();err != nil {
		t.Error(err)
	}else{
		fmt.Println(string(v))
	}
}

// benchmark test
func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i:=0;i<b.N;i++{
		err := json.Unmarshal([]byte(jsonStr),e)

		if err != nil{
			b.Error(err)
		}

		if _,err = json.Marshal(e);err !=nil {
			b.Error(err)
		}
	}
}


func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()

	e := Employee{}
	for i:=0;i< b.N;i++{
		err :=e.UnmarshalJSON([]byte(jsonStr))
		if err != nil{
			b.Error(err)
		}

		if _,err = e.MarshalJSON();err!=nil{
			b.Error(err)
		}
	}
}


