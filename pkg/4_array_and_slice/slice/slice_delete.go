package main

import (
	"fmt"
	"slices"
)

// 切片的删除
func SliceDelete() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("初始切片", slice)
	// 通过移动数据指针以删除开头的 X 个元素
	fmt.Println("删除前1个元素", slice[1:])
	fmt.Println("删除前3个元素", slice[3:])
	// 也可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成
	// （所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
	// ！！！！用白话说：这种方式会影响原始变量，要万分注意！！！
	// fmt.Println(append(slice[:0], slice[2:]...)) // 删除开头2个元素

	fmt.Println(slice)

	// 通过移动数据指针以删除末尾的 X 个元素
	fmt.Println("删除后1个元素", slice[:len(slice)-1])
	fmt.Println("删除后3个元素", slice[:len(slice)-3])
}

// 删除指定的元素
func SliceDeleteSpecifiedEle() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g"}

	needDelEle := 2

	fmt.Println(slice[:needDelEle], slice[needDelEle+1:])

	newSlice := append(slice[:needDelEle], slice[needDelEle+1:]...)

	fmt.Println(newSlice)
}

// 使用 1.21+ 版本之后新增的 Delete 方法删除元素
func SlicesDeleteElement() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g"}
	needDelEle := 2
	fmt.Println(slice[:needDelEle], slice[needDelEle+1:])
	// Delete() 方法会返回一个新的切片，原切片不会被修改
	// Delete() 参数：
	// 1. s：要删除元素的切片
	// 2. i：要删除的元素的起始索引号（包含）
	// 3. j：要删除的元素的结束索引号（不包含）
	// i.e. 第 i 个元素会删除，第 j 个元素不会删除
	newSlice := slices.Delete(slice, needDelEle, needDelEle+1)
	fmt.Println(newSlice)
}
