// 排序
package main

import (
	"fmt"
	"sort"
)

func main() {
	// 排序方法试针对内置数据类型的
	// 排序是原地更新，所以会改变给定序列的原始顺序且不会返回一个新值
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 这是一个int排序的例子
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// 我们也可以使用sort包来检查一个序列是不是已经排好序的。若已经排序好了，则输出布尔类型的结果
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
