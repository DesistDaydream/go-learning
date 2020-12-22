package main

import (
	"fmt"
)

// FormulaOperator 定义了一个接口，该接口包含一个求面积的方法和一个求周长的方法
type FormulaOperator interface {
	Area() float32
	Perimeter() float32
}

// 在 pkg/ 下定义了两个结构体及其方法，结构体实现了接口FormulaOperator
// 名为正方形的结构体，里面有一个属性是边长。名为矩形的结构体里面有两个属性是长和宽

// PrintResult 输出计算结果
func PrintResult(shape string, result FormulaOperator) {
	fmt.Printf("%s的面积：%f\n", shape, result.Area())
	fmt.Printf("%s的周长：%f\n", shape, result.Perimeter())
}
