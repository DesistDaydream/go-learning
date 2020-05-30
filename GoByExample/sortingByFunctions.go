// 使用函数自定义排序
// 有时候我们想使用和集合的自然排序不同的方法对集合进行排序
// 比如这篇代码中想按照字母的长度，而不是首字母顺序对字符串进行排序
package main

import (
	"fmt"
	"sort"
)

// 创建一个内置类型的别名
type ByLength []string

// 自定义类型实现了3个方法
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banban", "kiwi"}
	// 通过将原始的切片转换成ByLength来实现自定义排序
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
