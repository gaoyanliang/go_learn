package pipe_filter

import (
	"fmt"
	"reflect"
)

type StraightPipeline struct {
	Name string
	Filters *[]Filter
}

func NewStraightPipeline(name string,filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:name,
		Filters:&filters,
	}

}

func (f *StraightPipeline) Process(data Request) (Response,error){
	var ret interface{}
	var err error
	for _,filter := range *f.Filters {
		ret ,err = filter.Process(data)
		if err != nil {
			fmt.Println(reflect.TypeOf(filter), "process failed. err is", err)
			return ret,err
		}
		fmt.Println(reflect.TypeOf(filter), "process successed. ret is", ret)
		// 前面一个 filter 处理的 重新赋值给data
		data = ret
	}

	return ret,nil
}