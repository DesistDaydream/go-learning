package main

import "fmt"

func f() {
	for i := 0; i < 4; i++ {
		//此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

func main() {
	// f()

	// 一个匿名函数，实现自身逻辑，不与外界通信
	func(x, y int) {
		fmt.Println(x + y)
	}(3, 4)
	func(x, y int) { fmt.Println(x + y) }(3, 4)
}
