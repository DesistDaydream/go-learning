package main

import (
	"fmt"
)

// 定义一个名为Struct1的结构体，里面有3个字段
type structDemo struct {
	i1  int
	f1  float32
	str string
}

func structDemoFunction() {
	// 声明一个名为 sd 的变量，使用 structDemo 类型,也可以使用 new 函数进行初始化
	var sd structDemo
	// sd := new(structDemo)

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
