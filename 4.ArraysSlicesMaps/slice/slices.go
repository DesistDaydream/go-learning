package main

import "fmt"

func main() {
	fmt.Println("1.切片的基本用法")
	// slices()

	fmt.Println("\n2.切片的重组示例")
	reslice()

	fmt.Println("\n3.切片的追加append示例")
	// append1()

	fmt.Println("\n4.切片的复制copy示例")
	// copy1()

	fmt.Println("\n5.字符串、数组和切片的应用")
	// sliceString()

}

// 切片的基本用法
func slices() { 
	var arr1 = [7]int{2, 4, 6, 8, 10, 12, 14}
	// 初始化切片slice1，从数组arr1中引用索引号为2,3,4的元素
	var slice1 []int = arr1[2:5]

	// 打印切片slice1的所有元素的值，并打印切片的长度和容量
	// For-Range结构应用于数组和切片并迭代其所有的值，并返回切片的每个索引i与对应的值v
	for i, v := range slice1 {
		fmt.Printf("切片slice1的索引%d的值为%d\n", i, v)
	}

	fmt.Printf("这个arr1数组的长度%d\n", len(arr1))
	fmt.Printf("这个slice1切片的长度%d\n", len(slice1))
	// 切片的容量是从声明切片时所引用的对应数组的start索引开始一直到最后
	fmt.Printf("这个slice1的容量为%d\n", cap(slice1))

	// 增加切片的长度并打印所有切片元素
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	// slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range
}

// 切片重组(reslice)
func reslice() {
	slice1 := make([]int, 0, 10)
	//读取这个切片, cap(slice1)指的是slice1的容量，容量为10
	//每一次循环，切片slice1的长度就+1
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}
	//打印这个切片中的所有元素
	//上面的循环已经把len(slice1)长度增加到10
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}

//切片的追加(append)
func append1() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

// 切片的复制
func copy1() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	fmt.Println(slice1 ,slice2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

// 字符串、数组和切片的应用
func sliceString() {
	s := "hello"
	c := []byte(s)
	fmt.Println("使用byte函数把字符变量的所有字符转变成byte", c)
	c[0] = 'c'
	s2 := string(c) // s2 == "cello"
	fmt.Println("变更后的数组c为：", c)
	fmt.Println("变更后的字符串为:", s2)
}