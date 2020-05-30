package main

import (
	"fmt"
)

// 在函数体内直接对变量赋值，不会改变外部引用进来的变量的值
func zeroval(ival int) {
	ival = 0
}

// 该函数用了一个整型指针参数，在函数体内`解引用`该指针
// 从该内存地址得到该地址内对应的值。
// 对一个解引用的指针赋值会改变这个指针地址内的值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)
	// 通过`&i`语法来取得i的内存地址，i.e.指向i的指针
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	// 指针可以被打印，指针也是值的一种
	fmt.Println("pointer:", &i)
}
