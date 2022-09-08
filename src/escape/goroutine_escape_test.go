// goroutine 逃逸案例
package escape

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"testing"
	"time"
)

func queryAll() int {
	ch := make(chan int)
	// 向 channel 中写入 3 个查询结果
	for i := 0; i < 3; i++ {
		go func() {ch <- query()}()
	}
	return <-ch
}

func query() int {
	// 模拟查询耗时
	t := rand.Intn(100)
	time.Sleep(time.Duration(t) * time.Millisecond)
	return t
}

// 案例一：channel 发送不接收
func TestEscape1(t *testing.T) {
	for i := 0; i < 10; i++ {
		queryAll()
		fmt.Println("num of goroutine: ", runtime.NumGoroutine())
	}
}

// -----------------------

func recv_not_send() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan struct{}
	go func() {
		ch <- struct{}{}
	}()

	time.Sleep(time.Second)
}

// 逃逸案例二：接收不发送
func TestEscape2(t *testing.T) {
	recv_not_send()
}


// ---------------------------

func nil_chan() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var ch chan int
	go func() {
		ch <- 0
	}()

	time.Sleep(time.Second)
}

// 逃逸案例三： 空 channel
func TestEscape3(t *testing.T) {
	nil_chan()
}

// -------------------------------


func wait_slowly() {
	httpClient := http.Client{}
	//httpClient := http.Client{
	//	Timeout: time.Second * 15,
	//}

	for {
		go func() {
			_, err := httpClient.Get("https://xxx.com")
			if err != nil {
				fmt.Printf("http.Get err: %v\n", err)
			}
			// do something
		}()

		time.Sleep(time.Second)
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}
}

// 逃逸案例四：慢等待
func TestEscape4(t *testing.T) {
	wait_slowly()
}

// -----------------------

func not_unlock() {
	var total int = 0
	defer func() {
		fmt.Println("total: ", total)
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			// defer mutex.UnLock()
			total += 1
		}()
	}
}

// 逃逸案例五：加锁不解锁
func TestEscape5(t *testing.T)  {
	not_unlock()
}

// -----------------------------

func wait_group_escape(t int) {
	var wg sync.WaitGroup

	wg.Add(5)

	for i := 0; i < t; i++ {
		fmt.Println("wait group done")
		wg.Done()
	}
	wg.Wait()
}

func wait_group(t int) {
	var wg sync.WaitGroup
	for i := 0; i < t; i++ {
		wg.Add(1)
		defer wg.Done()
		fmt.Println("wait group done")
	}
	wg.Wait()
}

// 逃逸案例六： 同步锁使用不当
func TestEscap6(t *testing.T)  {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	//go wait_group_escape(3)
	go wait_group(10)
	time.Sleep(time.Second)
}

// -----------------------------

func TestEscape7(t *testing.T)  {
	//go http.ListenAndServe("0.0.0.0:6060", nil) //注册pprof监控
	ch := make(chan int, 5)                     //go channel，负责go routine之间通信，5个缓存
	flag := false                               //bool类型标记，模拟配置信息，false表示没有配置
	//第一个 go routine, 模拟写数据库信息，这里简化，直接读取channel的内容
	go func() {
		//当flag为false时，直接return，这行代码是导致go routine泄漏的关键
		if !flag {
			return
		}

		for {
			select {
			case recv := <-ch:
				//读取 channel
				fmt.Println("recive channel message:", recv)
			}
		}
	}()

	//for 循环模拟TCP长连接，每隔500ms向channel写一条数据，模拟心跳上报客户端自监控信息
	for {
		i := 0
		time.Sleep(500 * time.Millisecond)
		go func() {
			fmt.Println("goroutine count:", runtime.NumGoroutine())
			i++
			ch <- i
		}()
	}
}
