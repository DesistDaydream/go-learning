package main

import "fmt"

// 声明了两种数据类型，虽然底层类型都一样是float64，但是这两种新声明的是不同的数据类型
// 这两个新的数据类型不可以被互相比较或混在一个表达式运算。
// 刻意区分类型，可以避免一些像无意中使用不同单位的温度混合计算导致的错误；
// 因此需要一个类似Celsius(t)或Fahrenheit(t)形式的显式转型操作才能将float64转为对应的类型。
// Celsius(t)和Fahrenheit(t)是类型转换操作，它们并不是函数调用。类型转换不会改变值本身，但是会使它们的语义发生变化。
// 另一方面，CToF和FToC两个函数则是对不同温度单位下的温度进行换算，它们会返回不同的值。
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Println(CToF(10))
	fmt.Println(FToC(10))
	c := FToC(212.0)
	// fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
}
