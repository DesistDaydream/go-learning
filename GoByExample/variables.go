package main

import (
	"fmt"
)

func main() {
	// 用关键字`var`声明1个或多个变量
	var a string = "initial"
	fmt.Println(a)
	// 可以一次性声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)
	// 可以不指定变量类型，由Go自动推断已经初始化的变量类型
	var d = true
	fmt.Println(d)
	// 声明变量且没有给出对应的初始值时，变量将会初始化为`零`值。e.g.`int`类型的变量`零`值是`0`
	var e int
	fmt.Println(e)
	// `:=`语句是声明并初始化变量的简写。
	f := "short"
	fmt.Println(f)
}
