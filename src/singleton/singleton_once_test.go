/* 利用 sync.Once（func） 仅执行一次的特性实现 Singleton */
package singleton

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {}

var singletonObj *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Singletion Obj")
		singletonObj = new(Singleton)
	})
	return singletonObj;
}

func TestSingletion(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i ++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%d\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}