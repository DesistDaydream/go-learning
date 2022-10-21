package main

import (
	"fmt"
	"sort"
)

func SortDemo() {
	// 对字符串类型的切片进行排序
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	// ！！注意：这种排序是会改变原始变量中的值，而不是生成新额变量。！！
	fmt.Println("字符串切片排序:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("整型切片排序:  ", ints)

	// 检查一个切片是否已经排序
	fmt.Println("切片是否已经排序: ", sort.IntsAreSorted(ints))
}

func main() {
	SortDemo()
}
