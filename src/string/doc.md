## 字符串

1. string 是数据类型，不是引用或指针类型
2. string 是只读的 byte slice，len 函数可以返回他所包含的 byte 数（注意：不是字符个数 ）
3. string 的 byte 数组可以存放任何数据

常用字符串函数

- strings 包（https://golang.org/pkg/strings）
- strconv 包（https://golang.org/pkg/strconv）

## Unicode UTF8

1. Unicode 是一种字符集 （code point）
2. UTF8 是 unicode 的存储实现 （转换为字节序列的规则）

### 编码与存储

字符              "中"
Unicode          0x4E2D
UTF-8            0xE4B8AD
string/[]byte    [0xE4, 0xB8, 0xAD]

[Show me the code](string_test.go)