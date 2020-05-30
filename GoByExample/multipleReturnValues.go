package main

import (
	"fmt"
)

// `(int, int)`在函数中标志着这个函数返回2个int类型的值
func vals() (int, int) {
	return 3, 7
}

func main() {
	// 通过`多赋值`操作来使用这两个不通的返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	// 如果只需要函数返回值的一部分，可以使用`_`空白标识符
	_, c := vals()
	fmt.Println(c)
}
