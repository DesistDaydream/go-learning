package main

import "fmt"

func main() {
	// 通过 var 关键字声明变量
	var i int
	fmt.Println(&i, i) // 系统会为 var 声明的变量分配内存空间，并为变量赋予其类型的零值
}
