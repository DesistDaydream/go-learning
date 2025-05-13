package main

import "fmt"

// 泛型
// 使用类型形参编写 Go 函数以处理多种类型
// comparable 是一个内置 Constraint(约束)，用来表示类型形参可以接收的实参的种类，所谓的“约束”就是指，T 被约束为可以使用哪几种类型。
// comparable 包含所有可以比较类型，包括：booleans、numbers、strings、pointers、channels、可比较的 arrays、structs 中的属性
// comparable 可以改为 any，表示 T 可以是任意类型
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// 这里的 v 和 x 都是 T 类型
		// 而具体的 T 的类型由上层调用来决定，也就是**类型参数**
		// 若上层传递的参数是 string，则 s 和 x 也是 string；若上层传递的是 int，则 s 和 x 也是 int。
		if v == x {
			return i
		}
	}
	return -1
}

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

func GenericDemo[T Number](s, x T) T {
	return s + x
}

func main() {
	// 泛型。是对类型的抽象（i.e. 同一系列的对象群体，这个对象是指类型，这一类类型的抽象）（并不是对代码逻辑的抽象）
	// 在一个代码块中，若其处理逻辑对于不同的数据类型都是一样，这就是使用泛型的最佳时机之一
	// 非人话描述：对类型的抽象的相同逻辑泛用化
	fmt.Println(GenericDemo(1, 2))
	fmt.Println(GenericDemo("desist", "daydaream"))

	// Index() 函数适用于 int 类型的切片
	// 可以像这样 `Index[int](si, 15)` 传递参数告诉 `Index()` 当前应该使用哪种类型类型执行其中的代码，这里的 `[int]` 就是告诉 Index 的 T 应该是 int 类型，
	// 其调用相当于 `func Index[T comparable](s []T, x T) int {}` 变成了 `func Index(s []int, x int) int {}`
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	//  Index() 函数同样也使用于 string 类型的切片
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
