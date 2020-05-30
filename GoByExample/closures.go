package main

import (
	"fmt"
)

// 该函数的返回值为一个`函数`。
// 该返回的函数为在此函数体内定义的`匿名函数`
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// 调用函数，将其返回值(一个函数)赋值给一个变量。
	// 该变量中的函数的值包含了自己的值`i`，每次调用时都会更新`i`的值
	nextInt := intSeq()

	// 多次调用，系统不会自动注销函数，实现了数据持久化
	fmt.Printf("变量类型为：%T,变量的值为：%v\n", nextInt, nextInt())
	fmt.Printf("变量类型为：%T,变量的值为：%v\n", nextInt, nextInt())
	fmt.Printf("变量类型为：%T,变量的值为：%v\n", nextInt, nextInt())
	// 新变量，`i`值还原
	newInts := intSeq()
	fmt.Printf("变量类型为：%T,变量的值为：%v\n", nextInt, newInts())
}
