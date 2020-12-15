// 数组
package main

import (
	"fmt"
)

func arrays() {
	// 创建数组`a`，存放5个int类型的数据。元素的类型和氚度都是数组类型的一部分
	// 数组默认是`零`值，对于int类型数组来说，默认为`0`
	var a [5]int
	fmt.Println("emp:", a)
	// 使用`ArrayID[INDEX]=VALUE`语法来设置数组指定索引位置的值
	// 使用`ArrayID[INDEX]`得到指定索引位置的值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	// 使用内置函数`len`来获得数组的长度
	fmt.Println("len:", len(a))
	// 使用这个语法在一行内声明并初始化一个数组。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// 数组类型是一维的，但是你可以组合 构造多维的数据结构。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
