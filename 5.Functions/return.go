package main

import "fmt"

var num int = 10
var numx2, numx3 int

// Parameter 定义一个函数，并声明形参i和o，并指明类型为int，未定义返回值
func Parameter(i int, o int) {
	fmt.Println(i)
	fmt.Println(o)
}

// MultiPly3Nums 定义一个函数，计算3个参数相乘的结果，指定了3个形参和1个返回值
func MultiPly3Nums(a int, b int, c int) int {
	return a * b * c
}

// ThreeValues 定义一个函数，省略形参(注意加上括号，否则第一个括号内就编程该函数的参数而不是返回值了)，直接声明返回值的类型
func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}

// PrintValues 与后面两个函数组合使用，用于输出后面两个函数的返回值
//后面两个函数指的是getX2AndX3和getX2AndX3_2
//主要用于说明返回值是否指明具体的变量名
func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

// getX2AndX3 定义函数，指定参数input的类型为int，指定两个无名的int类型的返回值
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// getX2AndX3_2 定义函数，指定参数input的类型为int，还有两个有名的int类型的返回值。使用命名返回值的话，就算只有一个也要加上小括号
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	//如果是命名的返回值，则只需要一条简单的不带参数的return语句，否则想上一个函数一样，return需要带参数
	return
}

// MinMax 定义函数，比较两个整数的值，并返回最大值和最小值
func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else { // a = b or a < b
		min = b
		max = a
	}
	return
}

func main() {
	//调用函数并在调用时给函数传递实参，并把实参的值赋给形参
	Parameter(100, 200)

	//在函数中再调用函数，并把传递实参后被调用函数的返回值当成调用函数的实参
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))

	//调用函数ThreeValues，并使用空白符把其中一个返回值丢弃
	var i1 int
	var f1 float32
	i1, _, f1 = ThreeValues()
	fmt.Printf("The int: %d, the float: %f \n", i1, f1)

	//调用函数getX2AndX3和getX2AndX3_2，并把实参num带入
	//并把函数的返回值赋值给numx2和numx3，然后输出结果
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()

	//调用MinMax函数，输出最大值与最小值
	var min, max int
	min, max = MinMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)
}
