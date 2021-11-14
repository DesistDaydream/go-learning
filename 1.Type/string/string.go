package main

import "fmt"

func readOnlyString() {
	character := "Hello World"
	// 打印字符串的第二个字符
	fmt.Printf("%c\n", character[1])

	// 无法修改字符串中的一个字符，因为 Go 编译器会把上述定义的字符串内容分配到只读内存段
	// 下面这种写法是错误的
	// character[1] = 'a'
	// 如果非要修改字符串中的内容，可以将字符串转换为 Slices，Slices 概念详见单独章节，示例如下
	slicesCharacter := []byte(character)
	slicesCharacter[1] = 'E'
	fmt.Printf("%c\n", slicesCharacter)
}

func simpleString() {
	// 直接打印出字符串
	fmt.Println("Hello World")
	// 获取字符串的，结果为10进制，11，共11个字符，包括空格
	fmt.Println(len("Hello World"))
	// 打印字符串的第二个字符，用二进制表示，e在ASCII中的二进制为101
	fmt.Println("Hello World"[1])
	// 打印两段字符串加一起
	fmt.Println("Hello " + "World")
}

func main() {
	// simpleString()
	readOnlyString()
}
