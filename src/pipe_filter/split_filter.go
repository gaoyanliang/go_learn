package pipe_filter

import (
	"errors"
	"fmt"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string  // 分隔符
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

// Process 分隔字符串
func (sf *SplitFilter) Process(data Request)(Response,error){
	// 检查数据格式/类型，是否可以处理
	str,ok := data.(string)

	if !ok {
		return nil,SplitFilterWrongFormatError
	}
	parts := strings.Split(str,sf.delimiter)
	fmt.Print("SplitFilter Process ret: ")
	fmt.Println(parts)
	return parts,nil
}