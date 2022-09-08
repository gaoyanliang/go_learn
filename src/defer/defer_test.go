package _defer

/**
深入 Go 语言 defer 实现原理
https://www.luozhiyun.com/archives/523
 */
import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestDefer(t *testing.T)  {
	i := 1
	for {
		if i == 10 {
			break
		}
		defer fmt.Println(i)
		i ++
	}

	// -------------------------------
	v1 := f1(7)
	fmt.Println("v1 =", v1)

	v2 := f2()
	fmt.Println("v2 =", v2)

	v3 := f3()
	fmt.Println("v3 =", v3)

	v4 := f4()
	fmt.Println("v4 =", v4)

	// -------------------------------
	e1()
	e2()
	e3()

	// -------------------------------
	// 闭包引用了 x 变量，a b 可以看作是两个不同的实例，实例之间互不影响，同一个实例内 x 变量是同一个地址，因此具有累加效应
	var a = Accumulator()
	fmt.Println("value a: ")
	fmt.Printf("%d\n", a(1))
	fmt.Printf("%d\n", a(10))
	fmt.Printf("%d\n", a(100))

	var b = Accumulator()
	fmt.Println("value b: ")
	fmt.Printf("%d\n", b(1))
	fmt.Printf("%d\n", b(10))
	fmt.Printf("%d\n", b(100))


	// -------------------------
	deferCall()

	// -------------------------
	deferFor()
}

/**
defer 可以修改有名返回值函数的返回值
注意：只能修改有名返回值(named result parameters)函数，匿名返回值函数是无法修改的，如 f2
 */
// f returns arg * 6
func f1(arg int) (ret int) {
	defer func() { ret *= 6}()
	return arg
}

/**
匿名返回值函数是在return执行时被声明，因此在defer语句中只能访问有名返回值函数，而不能直接访问匿名返回值函数。
f2 return 100
 */
func f2() int {
	i := 100
	defer func() { i = i + 5 }()
	return i
}

// f3 return 1
func f3() (r int) {
	// 值传递，defer func 方法内部修改了 r，方法外部 r 没有修改
	defer func(r int) { r = r + 5}(r)
	return 1
}

// f4 return 6
func f4() (r int) {
	defer func(r *int) { *r = *r + 5}(&r)
	return 1
}

func e1()  {
	var err error

	defer fmt.Println("e1 =", err)

	err = errors.New("defer one error")
}

func e2() {
	var err error

	defer func() {
		fmt.Println("e2 =", err)
	}()

	err = errors.New("defer two error")
}

func e3() {
	var err error

	defer func(err error) {
		fmt.Println("e3 =", err)
	}(err)

	err = errors.New("defer three error")
}

// --------
// 返回类型为 func(int) int
func Accumulator() func(int) int {
	var x int

	return func(delta int) int {
		fmt.Printf("x's address: %+v, x's value: %+v - ", &x, x)
		x += delta
		return x
	}
}

// -----------
/**
defer here
defer caller
recover success. err:  should set user env.
 */
func deferCall()  {
	defer fmt.Println("===> deferCall start")
	var user = ""

	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err: ", err)
			}
		}()

		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic(errors.New("should set user env."))
			}

			// 此处不会执行
			fmt.Println("after painc")
		}()
	}()
	time.Sleep(2000)

	fmt.Println("===> end of deferCall")
}

func deferFor() {
	// 43210
	for i := 0; i < 5; i ++ {
		defer fmt.Println(i, 1)
	}

	// 55555
	for i := 0; i < 5; i ++ {
		defer func() {
			fmt.Println(i, 2)
		}()
	}

	// 55555
	for i := 0; i < 5; i ++ {
		defer func() {
			j := i
			fmt.Println(j, 3)
			// fmt.Println(j, &j, &i, 3)
		}()
	}

	// 43210
	for i := 0; i < 5; i ++ {
		j := i
		defer fmt.Println(j, 4)
	}

	// 43210
	for i := 0; i < 5; i ++ {
		j := i
		defer func() {
			fmt.Println(j, 5)
		}()
	}

	// 43210
	for i := 0; i < 5; i ++ {
		defer func(j int) {
			fmt.Println(j, 6)
		}(i)
	}
}