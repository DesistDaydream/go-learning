package main

import (
	"fmt"
)

// Struct(结构体) 是一组 Field(字段) 的集合，是一种复合数据类型。
// 每个字段都可以是任意类型的数据。
// **结构体的定义**。定义一个名为 structDemo 的结构体，里面有3个字段
type structDemo struct {
	i1  int
	f1  float32
	str string
}

func newStructDemo() *structDemo {
	return &structDemo{}
}

func structDemoFunction() {
	// **结构体的声明**，声明一个名为 sd 的 structDemo 类型的变量
	var sd *structDemo
	fmt.Printf("%p\n", sd)

	// **结构体的实例化**
	// 可以使用 new() 函数进行实例化
	// sd := new(structDemo)
	// 当然，最常见的是自己定义一个函数用以实例化结构体，类似于自己实现一个功能更全面的 new 函数
	// sd := newStructDemo()

	// 对结构体中的每个字段进行赋值
	sd.i1 = 10
	sd.f1 = 15.5
	sd.str = "Chris"
	// 也可以直接使用下面的初始化方式，对结构体进行初始化赋值，这种赋值方式是按照字段从上到下的顺序进行赋值
	// sd := &structDemo{10, 15.5, "Chris"}

	fmt.Printf("The int is: %d\n", sd.i1)
	fmt.Printf("The float is: %f\n", sd.f1)
	fmt.Printf("The string is: %s\n", sd.str)
	fmt.Println(sd)
}

func main() {
	structDemoFunction()

}
