/*
Benchmark 测试
func BenchmarkConcatStringByAdd(b *testing.B) {
	// 与性能测试无关的代码
	b.ResetTimer()
	for i := 0; i < b.N; i ++ {
		// 测试代码
	}
	b.StopTimer()
	// 与性能测试无关的代码
}

go test -bench=. -benchmen

- bench = <相关 benchmark 测试>
- Windows 下使用 go test 命令时，-bench=. 应该写为 -bench="."

 */
package benchmark

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByAdd(t *testing.T){
	assert := assert.New(t)
	elems := []string{"1","2","3","4","5"}
	ret := ""

	for _,elems := range elems {
		ret += elems
	}
	assert.Equal("12345",ret)
}


func TestConcatStringByByteBuffer(t *testing.T){
	assert := assert.New(t)
	var buf bytes.Buffer
	elems :=[]string{"1","2","3","4","5"}

	for _,elem := range elems {
		buf.WriteString(elem)
	}

	assert.Equal("12345",buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1","2","3","4","5"}

	b.ResetTimer()

	for i:=0 ; i< b.N;i++{
		ret := ""
		for _,elems := range elems {
			ret += elems
		}
	}

	b.StopTimer()
}


func BenchmarkConcatStringByByteBuffer(b *testing.B) {
	elems :=[]string{"1","2","3","4","5"}
	b.ResetTimer()

	for i:=0;i < b.N;i++ {
		var buf bytes.Buffer
		for _,elem := range elems {
			buf.WriteString(elem)
		}
	}

	b.StopTimer()
}
