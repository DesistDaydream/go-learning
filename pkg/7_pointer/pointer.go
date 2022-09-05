package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var (
	a int
)

func BasePointer() {
	// 输出a的值与a的地址值
	fmt.Printf("变量a的值为:%v\t变量a的指针所指向的地址值为：%v\n", a, &a)

	// 指针必须要有一个合法指向的变量的地址才可以操作，否则会报错
	// 定义一个指针类型的变量p，并赋值为空，但是会报错。因为指针p没有指向任何的地址
	// var p *int	//单独定义指针，值为nil，会报错
	// p = &a	//在p有合法指向的时候，p即可使用
	// 使用new函数来定义一个指针，会为该指针分配一个没有标识符的变量的地址
	// 这种用法叫做动态分配空间
	p := new(int)
	fmt.Println("指针变量p所指向的地址值为：", p, "该指针p所指向的无标识符变量的值为:", *p)
	*p = 6
	a = *p
	fmt.Println(a, *p, &a, p)

	// 把变量a的地址值给p，两者的地址值是不一样的
	// 所以指针指向的是变量地址，且指针的指针还是一个指针
	p = &a
	fmt.Printf("变量p的类型为：%v\n", reflect.TypeOf(p))
	fmt.Printf("变量p的值为：%v\n变量a的值为%v\n变量p的指针地址为：%v\n变量a的指针地址为：%v\n", p, a, &p, &a)
}

func main() {
	normalVar := 5
	// 通过 & 符号取得变量 a 的内存地址，即指向 a 的指针
	fmt.Println("变量 a 的内存地址，即指针为：", &normalVar)

	// 声明一个指针类型的变量
	// TODO: 使用 var 和 new() 有什么区别？ 使用 var 的话，没法 *ptr="a" 这么用，但是用 new() 就可以
	var ptr *string // 这是一个字符串指针类型的变量
	// ptr := new(string)
	fmt.Println("刚刚声明的指针没有任何内存地址，默认值为 nil：", ptr)

	// 指针赋值
	// 错误示例：不可以使用 *int 类型给 *string 类型赋值，虽然都是内存地址，但是不可以这么做
	// ptr = &normalVar
	// 正确示例:
	newPtr := strconv.Itoa(normalVar)
	ptr = &newPtr
	fmt.Println("为指针类型变量赋予一个内存地址后，获取该内存地址内保存的值：", *ptr)
}
