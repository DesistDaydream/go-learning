package main

import (
	"fmt"
	"reflect"
)

// InterfaceDemo 定义了一个名为 InterfaceDemo 的接口,
type InterfaceDemo interface {
	ToDoSomthing()
}

// MyString1 定义了一个基础类型为字符串的名为MyString的类型
type MyString1 string

// ToDoSomthing 让 MyString实现了接口 InterfaceDemo
// 因为 MyString1 实现了 ToDoSomthing 方法，且 InterfaceDemo 接口中的所有方法该类型都实现了，所以 MyString1 实现了 InterfaceDemo 接口
func (ms MyString1) ToDoSomthing() {
	fmt.Println("我是让 MyString1 结构体实现 InterfaceDemo 这个接口的 ToDoSomthing 这个方法中，输出的内容")
}

// MyString2 is
type MyString2 string

// ToDoSomthing is
func (ms MyString2) ToDoSomthing() {
	fmt.Println("我是让 MyString2 结构体实现 InterfaceDemo 这个接口的 ToDoSomthing 这个方法中，输出的内容")
}

func main() {
	name1 := MyString1("我是 MyString1 类型声明的变量的值")
	name2 := MyString2("我是 MyString2 类型声明的变量的值")
	// 声明一个接口类型的变量(接口类型可以简称为 接口)
	var v InterfaceDemo
	// 接口已声明未使用的时候，接口变量的类型为nil(空),值也为nil(空)，接口的类型会随着值的不同而又不同的类型
	fmt.Printf("变量 name1 的类型为：%T, 变量 name2 的类型为：%T\n接口类型的变量 v 的类型为：%T; 值为：%v\n\n", name1, name2, v, v)

	// 由于 MyString1 和 MyString2 类型实现了 VowelsFinder 接口
	// 所以这些自定义类型的变量可以作为值，赋给接口类型的变量，且接口变量的类型变为实现该接口的自定义类型
	v = name1
	// 调用方法
	v.ToDoSomthing()
	fmt.Printf("自定义类型为:%v,接口类型为:%v\n\n", reflect.TypeOf(name1), reflect.TypeOf(v))

	v = name2
	// 调用方法
	v.ToDoSomthing()
	fmt.Printf("自定义类型为:%v,接口类型为:%v\n", reflect.TypeOf(name2), reflect.TypeOf(v))
}
