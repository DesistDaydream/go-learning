// 接口(interfaces)是命名了的方法签名(signatures)的集合
package main

import (
	"fmt"
	"math"
)

// 这是基本的接口
type geometry interface {
	area() float64
	perim() float64
}

// 将在类型`rect`和`circle`上实现这个接口
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 要想让一个类型实现一个接口，我们就需要实现接口中的所有方法
// 这是在类型`rect`上实现了`geometry`接口
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// 这是在`circle`类型上实现了接口
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量具有接口类型，则可以调用指定接口中的方法。
// 这个函数`measure`可以用来在任何的`geometry`上工作
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func interfaceType() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	// 结构体类型`circle`和`rect`都实现了`geometry`接口
	// 所以我们可以使用这些结构体的实例作为`measure`的参数(尽管这个形参是接口类型而不是结构体类型)。
	measure(r)
	measure(c)
}
