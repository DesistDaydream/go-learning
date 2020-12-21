package main

import (
	"github.com/DesistDaydream/GoLearning/7.Structs_Interfaces/interface/practice/pkg/rectangle"
	"github.com/DesistDaydream/GoLearning/7.Structs_Interfaces/interface/practice/pkg/square"
)

// Formula 计算不同图形的周长和面积
func Formula() {
	square := &square.Square{Side: 5}
	rectangle := rectangle.Rectangle{Length: 4, Width: 3}
	PrintResult("正方形", square)
	PrintResult("矩形", rectangle)
}

// main 中的函数调用相当于模拟其他 package 调用本 package 的情景
func main() {
	// 计算图形面积和周长练习
	Formula()

	// 计算工资练习
	Salary()
}
