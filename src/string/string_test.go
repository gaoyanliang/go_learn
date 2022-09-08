package string

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

/*string是一个数据类型，不是指针或者引用类型
  string的byte数组可以存放任意数据
*/
func TestStringAccess(t *testing.T) {
	var s string
	t.Logf("sting is *%s*", s)
	s = "hello"
	t.Logf("sting is *%s*", s)
	t.Logf("string is *%s*, and len is %d", s, len(s))
	//s[1] = "s" // 字符串是一个只读的 byte slice，不能被修改
	s = "\xE4\xB8\xA5"
	t.Logf("string is *%s*, and len is %d", s, len(s))
	t.Logf("%s 's unicode is %x", s, s) // e4b8a5

	s = "\xE4\xB8\xA5\xBB" // 错误的编码也可以输出
	t.Log(s, len(s))
}

func TestStringUnicodeAndUTF8(t *testing.T) {
	s := "中"
	t.Logf("string is *%s*, and len is %d", s, len(s))

	c := []rune(s)
	t.Logf("%s 's rune is %d, and len is %d", s, c, len(c))

	t.Logf("中`s unicode %x", c)
	t.Logf("中`s utf-8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人名共和国"
	for _, v := range s {
		// [1] 格式化
		t.Logf("%[1]c, %[1]x", v)
	}
}


// ----------------- 字符串函数 -----------------

// 1. 字符串分隔 & 连接
func TestStringFn(t *testing.T)  {
	s := "A,B,C"
	parts := strings.Split(s, ",")

	for _, p := range parts {
		t.Log(p)
	}
	
	t.Log(strings.Join(parts, "--"))
}

// 2. 字符串与整数的转换
func TestConv(t *testing.T) {
	a := 10
	s := strconv.Itoa(a);

	t.Logf("int is %d, conv string is [%s]", a, s)

	str := "10"
	if conv,err := strconv.Atoi(str); err == nil {
		t.Logf("10 + 10 = %d", a + conv)
	} else {
		t.Log("---compute error---")
	}
}

// 3. 获取字符串长度

func TestStringLen(t *testing.T) {
	s := "中华人名共和国"
	// len 字节长度：和编码无关
	fmt.Println("s len = ", len(s))
	// 字符数量，和编码有关
	fmt.Println("s length = ", utf8.RuneCountInString(s))

	s1 := "中华人名共和国--"
	// len 字节长度：和编码无关
	fmt.Println("s1 len = ", len(s1))
	// 字符数量，和编码有关
	fmt.Println("s1 length = ", utf8.RuneCountInString(s1))

}

