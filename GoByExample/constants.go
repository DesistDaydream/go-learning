package main

import (
	"fmt"
	"math"
)
// 关键字`const`用于声明一个常量
const s string = "constant"

func main() {
	fmt.Println(s)
	// `const`语句可以出现在任何`var`语句可以出现的地方
	const n = 500000000
	// 常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)
	// 数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	fmt.Println(int64(d))
	// 当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。
	// e.g.这里的`math.Sin`函数需要一个`float64`的参数。
	fmt.Println(math.Sin(n))
}
