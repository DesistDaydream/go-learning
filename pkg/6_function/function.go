package main

import "fmt"

func callByValue(passedValue int) {
	passedValue = 0
}

func callByReference(passedReference *int) {
	*passedReference = 0
}

func main() {
	// Go 通过指针，允许程序通过`引用传递`来传递值和数据结构
	i := 1
	fmt.Println("原始值:", i)

	// 通过值传递，在函数内修改不会影响原始值
	callByValue(i)
	fmt.Println("值传递:", i)

	// 通过引用传递，在函数内修改会影响原始值
	callByReference(&i)
	fmt.Println("引用传递:", i)
}
