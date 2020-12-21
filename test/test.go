package main

import "fmt"

// FormulaOperator 定义了一个接口，包含两个方法
// 该接口实现了一个求面积的方法和一个求周长的方法
// type FormulaOperator interface {
// 	Area() float32
// 	Perimeter() float32
// }

// 定义了两个结构体，结构体实现了接口FormulaOperator
// 名为正方形的结构体，里面有一个参数是边长。名为矩形的结构体里面有两个参数

// Square 正方形的属性
type Square struct {
	side float32
}

// Rectangle 长方形的属性
type Rectangle struct {
	length float32
	width  float32
}

// Area 正方形求面积的方法，接收了正方形的结构体并使用结构体中的边长属性来计算面积
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// Perimeter 正方形求周长的方法
func (sq *Square) Perimeter() float32 {
	return sq.side * 4
}

// Area 矩形求面积的方法
func (r Rectangle) Area() float32 {
	return r.length * r.width
}

// Perimeter 矩形求周长的方法
func (r Rectangle) Perimeter() float32 {
	return r.length*2 + r.width*2
}

// PrintResult 输出计算结果
func PrintResult(shape string, result *Square) {
	fmt.Printf("%s的面积：%.2f\n", shape, result.Area())
	fmt.Printf("%s的周长：%.2f\n", shape, result.Perimeter())
}

// PrintResult2 输出计算结果
func PrintResult2(shape string, result *Rectangle) {
	fmt.Printf("%s的面积：%.2f\n", shape, result.Area())
	fmt.Printf("%s的周长：%.2f\n", shape, result.Perimeter())
}

// Formula 计算不同图形的周长和面积
func main() {
	square := &Square{side: 5}
	rectangle := &Rectangle{length: 4, width: 3}
	PrintResult("正方形", square)
	PrintResult2("矩形", rectangle)

}
