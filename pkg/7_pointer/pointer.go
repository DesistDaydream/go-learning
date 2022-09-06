package main

import (
	"fmt"
	"strconv"
)

func BasePointer() {
	normalVar := 5
	// 通过 & 符号取得变量 a 的内存地址，即指向 a 的指针
	fmt.Println("变量 a 的内存地址，即指针为：", &normalVar)

	// 声明一个指针类型(字符串指针类型)的变量，两种方式：
	var ptrVar *string
	fmt.Println("通过 var 关键字声明的指针变量的值，默认为 nil：", ptrVar)
	fmt.Println("但是这个指针变量本身是存在内存地址的：", &ptrVar)
	// 由于此时 ptr 为 nil，并不具有内存地址，所以不能通过解引用的方式为指针所指向的内存地址中添加值。毕竟现在还没有内存地址了
	// *ptrVar = "pointer string"

	// 指针变量的赋值。
	// 给指针类型的变量赋值，赋予的是内存地址，不是字符串、整数之类的。内存地址类似于这样：0xc000014088
	newPtrVar := strconv.Itoa(normalVar)
	ptrVar = &newPtrVar
	fmt.Println("为指针变量赋予一个内存地址后，即可以通过该变量获取到内存地址中保存的值：", *ptrVar)
	// 错误示例：不可以使用 *int 类型给 *string 类型赋值，虽然指针变量的值都是内存地址，但是不可以这么做
	// ptr = &normalVar
}

func Instantiation() {
	// 我们可以使用 new() 函数直接实例化一个 *string 类型的变量，此时系统会自动为指针变量赋予一个值。
	// 也就是说，实例化一个指针类型的变量，系统会开辟出两块内存空间
	// 1. 第一块内存空间存储指针变量本身的值
	// 2. 第二块内存空间就是变量的值所指向的内存空间，用以存放我们想要的值
	ptrVar := new(string)
	fmt.Println("通过 new() 函数实例化指针，系统会为该变量的值赋予一个内存地址：", ptrVar)

	// 由于此时 ptrNew 具有内存地址，所以可以直接通过解引用的方式为 ptrNew 内存地址空间内赋予值
	*ptrVar = "pointer string"
	fmt.Println("通过解引用的方式，直接将值放到指针变量的值所指向的内存地址内部：", *ptrVar)
}

func main() {
	// BasePointer()
	Instantiation()
}
