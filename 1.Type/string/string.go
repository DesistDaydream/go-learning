package main

import "fmt"

func main() {
	//直接打印出字符串
	fmt.Println("Hello World")
	// 获取字符串的，结果为10进制，11，共11个字符，包括空格
	fmt.Println(len("Hello World"))
	//打印字符串的第二个字符，用二进制表示，e在ASCII中的二进制为101
	fmt.Println("Hello World"[1])
	//打印两段字符串加一起
	fmt.Println("Hello " + "World")
}