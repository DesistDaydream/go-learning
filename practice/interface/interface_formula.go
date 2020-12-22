package main

import (
	"fmt"
)

// FormulaOperator 定义了一个接口(用于进行公式运算)，该接口包含一个求面积的方法和一个求周长的方法
type FormulaOperator interface {
	Area() float32
	Perimeter() float32
}

// 在 formula/ 下定义了两个结构体及其方法，结构体实现了 FormulaOperator 接口

// PrintResult 输出计算结果
func PrintResult(shape string, result FormulaOperator) {
	fmt.Printf("%s的面积：%f\n", shape, result.Area())
	fmt.Printf("%s的周长：%f\n", shape, result.Perimeter())
}
