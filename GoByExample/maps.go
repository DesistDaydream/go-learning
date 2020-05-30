package main

import (
	"fmt"
)

func main() {
	// 使用make函数创建一个新的空map
	m := make(map[string]int)
	// 使用`MapID[KEY]=VALUE`语法来设置指定map中的键值对
	m["k1"] = 7
	m["k2"] = 13
	// 输出指定map的所有键值对
	fmt.Println("map:", m)
	// 使用`MapID[KEY]`来获取指定键的值
	v1 := m["k1"]
	fmt.Println("v1:", v1)
	// 使用len函数返回键值对的数目
	fmt.Println("len:", len(m))
	// 使用delete从指定的map中移除指定的键值对
	delete(m, "k2")
	fmt.Println("map:", m)
	// 当从一个 map 中取值时，可选的第二返回值指示这个键 是否在这个 map 中。
	// 这可以用来消除键不存在和键有零值， 像 0 或者 "" 而产生的歧义。这里我们不需要值，所以 用_空白标识符(blank identifier)_忽略。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 在一行声明和初始化一个新的map
	n1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n1)
	// 另外一种声明以及初始化map的方法
	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("map:", n)
}
