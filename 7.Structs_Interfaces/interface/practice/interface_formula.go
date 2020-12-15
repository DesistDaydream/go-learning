package main

import "fmt"

// FormulaOperator 定义了一个接口，包含两个方法
// 该接口实现了一个求面积的方法和一个求周长的方法
type FormulaOperator interface {
	Area() float32
	Perimeter() float32
}

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
func PrintResult(shape string, result FormulaOperator) {
	fmt.Println(result.Area())
	fmt.Println(result.Perimeter())
	fmt.Printf("这个%s的面积：%f\n", shape, result.Area())
	fmt.Printf("这个%s的周长：%f\n", shape, result.Perimeter())
}

// Formula 计算不同图形的周长和面积
func Formula() {
	square := &Square{side: 5}
	rectangle := Rectangle{length: 4, width: 3}
	PrintResult("正方形", square)
	PrintResult("矩形", rectangle)

}

/*
func Formula() {
	square := new(Square)
	square.side = 5
	// square := &Square{5} //另一种方式
	rectangle := Rectangle{4, 3}

	// 使用数组将结构体类型的变量赋值给一个接口类型的变量
	shapes := []Formula{square, rectangle}
	fmt.Println(square.Area())
	fmt.Println(square.Perimeter())
	fmt.Printf("这个正方形的面积：%f\n", shapes[0].Area())
	fmt.Printf("这个正方形的周长：%f\n", shapes[0].Perimeter())
	fmt.Println(rectangle.Area())
	fmt.Println(rectangle.Perimeter())
	fmt.Printf("这个矩形的面积：%f\n", shapes[1].Area())
	fmt.Printf("这个矩形的周长：%f\n", shapes[1].Perimeter())
	// 下面注释的几行也可以每次单独给接口赋值
	// var shapes Formula
	// shapes = square
	// fmt.Printf("这个正方形的面积：%f\n", shapes.Area())
	// fmt.Printf("这个正方形的周长：%f\n", shapes.Perimeter())
	// shapes = rectangle
	// fmt.Printf("这个矩形的面积：%f\n", shapes.Area())
	// fmt.Printf("这个矩形的周长：%f\n", shapes.Perimeter())
}
*/
// 不使用切片的情况下，直接给函数传参即可实现输出效果
