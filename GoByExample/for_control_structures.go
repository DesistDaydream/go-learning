package main

import (
	"fmt"
)

func forControl() {
	// 最基本的类型，具有单一条件的循环体
	// 如果i小于等于3，则执行循环体
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// 经典的for循环结构`初始化; 条件; 后续`
	// 后续是在`循环体`执行完成后执行的语句
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// 不带条件的for循环将一直执行
	// 直到在循环体内使用了`break`或者`return`来跳出循环
	for {
		fmt.Println("loop")
		break
	}
	// 使用`continue`来继续下一次循环而不执行continue关键字后面的语句
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
