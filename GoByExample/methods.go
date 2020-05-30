// 定义在某个类型上的函数叫方法
package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

// 这个`area`方法有一个`*rect类型`的`接收器`
func (r *rect) area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收器定义方法。以下是一个值类型的接收器
func (r rect) perim() int {
	return 2 * r.width + 2 * r.height
}
func main() {
	r := rect{width: 10, height: 5}
	// 调用上面为结构体定义的两个方法。
	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())
	// Go自动处理方法调用时的值和指针之间的转化。
	rp := &r
	fmt.Println("area", rp.area())
	fmt.Println("perim", rp.perim())
}
