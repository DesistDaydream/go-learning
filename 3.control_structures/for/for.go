package main

import (
	"fmt"
)

func baseFor() {
	// 基本for循环语句，初始化变量i的值为0，当i小于3时执行循环体，每执行一次i的值加1
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
}

func infiniteLoop() {
	i := 0
	for {
		fmt.Println(i)
		i += 1
		if i == 3 {
			return
		}
	}
}

func continueTest() {
	//continue用于忽略循环体内的continue关键字的后续代码块，而直接进入下一次循环的过程
	//continue只能被用于for循环中
	for i := 0; i < 10; i++ {
		//如果i等于5，那么则继续后面的循环，不执行当前循环的代码，i.e.不输出5的值
		if i == 5 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

func breakTest() {
	//break用于退出当前循环体，并继续执行后续代码,且只会退出最内层的循环
	//break可以用在多种结构体中
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			fmt.Printf("%d ", j)
		}
	}
}

func labelTest() {
	//标签，在标签名出现的地方，直接跳到标签定义时的位置继续执行代码
	//如果是`break LABEL1`的话，则会连最外层循环也退出,仅输出一行
LABEL1:
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				fmt.Printf("\n")
				continue LABEL1
			}
			fmt.Printf("%d,%d ", i, j)
		}
	}

	fmt.Println("goto的用法")
	//
	i := 0
HERE:
	fmt.Printf("%d ", i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

func forRange() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	//For-Range结构应用于数组和切片，index为索引，season为该索引位置元素的值
	for index1, season := range seasons {
		fmt.Printf("Season %d is: %s\n", index1, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

	for index1 := range seasons {
		fmt.Printf("%d ", index1)
	}
}

func main() {
	// 基本结构示例
	fmt.Println("1.基本for循环示例")
	baseFor()
	// 无限循环示例
	fmt.Println("\n2.无限循环示例")
	infiniteLoop()
	//continueTest功能示例
	fmt.Println("\n3.continue用法")
	continueTest()
	// break功能示例
	fmt.Println("\n4.break用法")
	breakTest()
	// label与goto用法示例
	fmt.Println("\n5.标签与goto的用法")
	labelTest()
	// for-range结构示例
	fmt.Println("\n6.for-range结构的用法")
	forRange()
}
