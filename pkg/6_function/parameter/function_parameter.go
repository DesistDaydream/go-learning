package main

import (
	"fmt"
)

func callback(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

// 定义函数，具有两个形参
// 第一个形参是整型
// 第二个形参是具有两个形参的函数。也就是回调函数
func callbackDemo(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func main() {
	// 调用 callbackDemo 函数，传递实参 1 和 callback
	// 实参 1 是整型；实参 callback 是一个函数，i.e. callback 是 callbackDemo 的回调函数
	callbackDemo(1, callback)
}
