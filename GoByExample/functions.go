package main

import (
	"fmt"
)

// 这是一个函数，接收两个int类型的参数，返回一个int类型的值
func plus(a int, b int) int {
	// Go需要`return`关键字明确的返回，不会自动返回最后一个表达式的值
	return a + b
}

// 当多个连续的参数为同样的类型时，最多可以仅声明最后一个参数类型，而忽略之前相同类型参数的类型声明
func plusPlus(a, b, c int) int {
	return a + b + c
}

// 通过FunctionID(ARGS)来调用函数
func functions() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
