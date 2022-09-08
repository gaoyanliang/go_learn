/*
通过 context 取消任务

- 根 context： 通过 context.BackGround() 创建
- 子 context： context.WithCancel(parentContext) 创建
ctx，cancel ：= context.WithCancel(context.Background())

当前 context 被取消时，基于他的子 context 都会被取消
通过 <-ctx.Done 接收取消通知
 */
package cancel_task_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/* 判断 context 是否完成 */
func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false

	}
}

func TestCancelByContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i ++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}


