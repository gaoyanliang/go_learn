package pipe_filter

import (
	"errors"
	"fmt"
	"reflect"
)

var SumFilterWrongFormatError = errors.New("input data should be []int")

type SumFilter struct {}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

// Process int 分片求和
func (sf SumFilter) Process(data Request) (Response, error) {
	fmt.Println("SumFilter received ", data, "type is ", reflect.TypeOf(data))

	elems, ok := data.([]int)
	if (!ok) {
		return nil, SumFilterWrongFormatError
	}

	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}

