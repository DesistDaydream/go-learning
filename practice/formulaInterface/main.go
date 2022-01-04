package main

import (
	"github.com/DesistDaydream/go-learning/practice/formulaInterface/formula"
	"github.com/DesistDaydream/go-learning/practice/formulaInterface/formulaInterface"
)

// main 中的函数调用相当于模拟其他 package 调用本 package 的情景
func main() {
	// 计算图形面积和周长练习
	formulaInterface.PrintResult("正方形", &formula.Square{Side: 5})
	formulaInterface.PrintResult("矩形", &formula.Rectangle{Length: 4, Width: 3})
}
