// Package pipe_filter 定义 filter 接口
package pipe_filter

// Request is the input of the filter
type Request interface {

}

// Response is the output of the filter
type Response interface {

}

type Filter interface {
	Process(data Request) (Response,error)
}
