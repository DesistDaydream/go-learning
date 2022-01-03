package main

import "fmt"

func main() {
	fmt.Println("1.测试for循环对数组的赋值以及引用输出")
	arrays()
	
	fmt.Println("\n2.测试数组的初始化")
	arrInit()

	fmt.Println("\n3.测试多维数组")
	multidim_array()
}

//数组的基本使用
func arrays() {
	//数组的声明。声明一个名为arr1，长度为5，类型为int的数组
	var arr1 [5]int
	//数组的赋值。使用for形式轮流对每个元素进行赋值
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}
	//遍历数组，并输出数组中的每一个元素
	for i := 0; i < len(arr1); i++ {
		//使用元素索引用每个元素并输出每个元素的值
		fmt.Printf("数组arr1索引为%d的值是%d\n", i, arr1[i])
	}

	fmt.Println("也可以直接整体输出一个数组：", arr1)

	fmt.Println("数组arr1的容量为：", cap(arr1))
}

//数组的初始化
func arrInit() {
	// 正常对数组中每个元素进行赋值
	// var arrAge = [5]int{18, 20, 15, 22, 16}

	// 初始化的时候不指定长度，直接赋值，赋值的个数就是长度
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrLazy = []int{5, 6, 7, 8, 22}

	// 字符型数组的初始化，直接对3,4号元素赋值，其余默认为空。数组长度为初始化时最大的元素号+1
	// var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	var arrKeyValue = []string{3: "Chris", 4: "Ron"}
	// For-Range结构应用于数组和切片并迭代其所有的值，并返回切片的每个索引i与对应的值v
	for i, v := range arrLazy {
		fmt.Printf("索引%d的值为%d\n", i, v)
	}

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}

func multidim_array() {
	const (
		WIDTH  = 2
		HEIGHT = 3
	)

	type pixel int
	var screen [WIDTH][HEIGHT]pixel

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}

	fmt.Println(screen)
}
