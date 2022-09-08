/*
sync.Pool 对象缓存

------------------------Processor------------------------
|                                                        |
|    ---------------            ---------------          |
|   ｜              ｜           |              |         |
|   ｜    私有对象   ｜           |     共享池    |         |
|    ---------------            ---------------          ｜
|       协程安全                     协程不安全             ｜
---------------------------------------------------------

对象获取

- 尝试从私有对象获取
- 私有对象不存在，尝试从当前 Processor 的共享池获取
- 如果当前 Processor 共享池也是空的，尝试去其他 Processor 的共享池中获取
- 如果所有子池都是空的，最后就用用户指定的 New 函数产生一个新的对象返回

对象放回

- 如果私有对象不存在，则保存为私有对象
- 如果私有对象存在，放入当前 Processor 子池的共享池中

pool := &sync.Pool {
	New: func() interface{} {
		return 0
	}
}
array := pool.Get().(int) // 断言，判断类型是否是 int
pool.Put(10)

sync.Pool 对象生命周期

- GC 会清理 sync.pool 缓存的对象
- 对象的缓存有效期为下一次 GC 之前

总结：

- 适用于通过复用，降低复杂对象的创建和GC代价
- 协程安全，会有锁的开销
- 生命周期受 GC 影响，不适合做连接池等，需要自己管理生命周期的资源池化

 */
package sync_pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T){
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object")
			return 100
			// return 10.1
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() // GC 会清除 sync.pool 中缓存的对象
	v1 , _ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolMultiGroutine(t *testing.T)  {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("Create a new object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i:=0; i <10  ; i++  {
		wg.Add(1)
		go func(i int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
