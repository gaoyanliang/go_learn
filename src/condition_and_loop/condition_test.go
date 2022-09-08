package condition_and_loop

import (
	"fmt"
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T)  {
	if a := 1 == 1; a {
		t.Log("1 == 1")
	}
}

//func TestIfMultiSec2(t *testing.T)  {
//	if v,err := someFun(); err == nil {
//		t.Log("success")
//	} else {
//		t.Log("fail")
//	}
//}

// 注意 go 语言中 不需要 break，默认存在break
func TestSwitch(t *testing.T)  {
	fmt.Printf("当前系统为 %s", runtime.GOOS)
	fmt.Println()

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("darwin")
		// break
		/* go is not need write break if program
		execute here will be auto exit or quit */
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}
}

func TestSwitch2(t *testing.T)  {
	num := 5
	switch  {
	case 0 <= num && num <= 3:
		fmt.Println("0-3")
	case 4 <= num && num <= 6:
		fmt.Println("4-6")
	case 7 <= num && num <= 9:
		fmt.Println("7-9")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i ++ {
	switch {
	//case 0, 2:
	case i % 2 == 0:
		t. Log("Even")
	//case 1, 3:
	case i % 2 == 1:
		t. Log ("Odd")
	default:
		t. Log("unknow")
	}
}
}