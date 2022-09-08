package client

import (
	/*
		包的导入依赖gopath，你的路径是/user/.../gohomework/go_learning
		所以从src下的路径开始导入
	*/
	"go_learn/src/package/series"
	"testing"
)

// 需要关闭 mod 模式 export GO111MODULE=off
func TestPackage(t *testing.T) {
	var (
		list []int
		err  error
	)
	if list, err = series.GetFibonacci(10); err != nil {
		t.Log(err)
		return
	}
	t.Log(list)
	t.Log(series.Square(3))

}

