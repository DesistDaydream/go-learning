package main

import "fmt"

// 泛型
// 使用类型形参编写 Go 函数以处理多种类型
// comparable 是一个内置 Constraint(约束)，用来表示类型形参可以接收的实参的种类，所谓的“约束”就是指，T 被约束为可以使用哪几种类型。
// comparable 包含所有可以比较类型，包括：booleans、numbers、strings、pointers、channels、可比较的 arrays、structs 中的属性
// comparable 可以改为 any，表示 T 可以是任意类型
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// 这里的 v 和 x 都是 T 类型，且 T 类型具有可以比较的约束，因此这里可以使用 ==
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index() 函数适用于 int 类型的切片
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	//  Index() 函数同样也使用于 string 类型的切片
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}