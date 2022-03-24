package main

import (
	"fmt"
)

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

//定义函数，具有两个形参
//第一个形参是整型
//第二个形参是具有两个形参的函数
func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func main() {
	//调用callback函数，传递实参1、Add
	//实参1是整型，实参Add是一个函数
	callback(1, Add)
}
