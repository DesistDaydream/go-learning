package main

import (
	"fmt"
)

func ifElse() {
	// 基本的if-else循环结构例子
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// 可以不要`else`关键字的if语句
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	// 在条件语句之前可以声明一个变量，该变量可以用于所有分支语句中
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
