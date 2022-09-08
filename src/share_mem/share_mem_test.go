package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestGoutine(t *testing.T) {
	counter :=0
	for i:= 0; i<5000;i++ {
		go func() {
			counter ++
		}()
	}
	time.Sleep(time.Millisecond*1)
	// counter ！= 5000 非并发安全
	t.Logf("counter= %d",counter)

}


func TestCounterThreadSafe(t *testing.T) {
	mut := sync.Mutex{}
	counter :=0
	for i:= 0; i<5000;i++ {
		go func() {
			// 解锁
			defer func() {
				mut.Unlock()
			}()
			// 内存锁
			mut.Lock()

			counter ++
		}()
	}
	// 等待协程
	time.Sleep(time.Millisecond*1)
	t.Logf("counter= %d",counter)

}


func TestCounterWaitGroup(t *testing.T) {
	mut := sync.Mutex{}
	var wg sync.WaitGroup
	counter :=0
	for i:= 0; i<5000;i++ {
		// 添加任务
		wg.Add(1)
		go func() {
			// 解锁
			defer func() {
				mut.Unlock()
			}()
			// 内存锁
			mut.Lock()

			counter ++
			// 任务结束
			wg.Done()
		}()
	}
	// 执行等待
	wg.Wait()
	t.Logf("counter= %d",counter)

}

