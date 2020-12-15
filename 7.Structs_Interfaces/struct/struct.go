package main

import "fmt"

func main() {
    struct2()
}

// 定义一个名为Struct1的结构体，里面有3个字段
type struct1 struct {
    i1  int
    f1  float32
    str string
}

func struct2() {
	// 声明一个名为ms的变量，使用struct1类型,也可以使用new函数进行初始化
	var ms struct1
	// ms := new(struct1)

	// 对结构体中的每个字段进行赋值
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Chris"
	// 也可以直接使用下面的初始化方式，对结构体进行初始化赋值，这种赋值方式是按照字段从上到下的顺序进行赋值
	// ms := &struct1{10, 15.5, "Chris"}

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)
}
