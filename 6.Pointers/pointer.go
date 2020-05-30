package main

import (
	"fmt"
	"reflect"
)

var (
	a int
)

func BasePointer() {
	a = 5
	// 输出a的值与a的地址值
	fmt.Printf("变量a的值为:%v\t变量a的指针所指向的地址值为：%v\n",a, &a)
	
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
	fmt.Printf("变量p的类型为：%v\n",reflect.TypeOf(p))
	fmt.Printf("变量p的值为：%v\n变量a的值为%v\n变量p的指针地址为：%v\n变量a的指针地址为：%v\n",p, a, &p, &a)
}

func doSomething(i *int, num int) {
	*i = num
	fmt.Println(i, num, *i)
}

func main() {
	BasePointer()
	// doSomething(&a, 10)
}