package main

import (
	"fmt"
)

func main() {
	// 使用range来遍历切片，对于数组也可以使用这种方式。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	// range会返回数组和切片中每个元素的`索引`和`值`的。
	// 上面我们不需要索引，所以我们使用 空白标识符`_`来忽略它。
	// 有时候我们实际上是需要这个索引的。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	// range在map中迭代键值对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// range也可以只遍历map的键
	for k := range kvs {
		fmt.Println("key:", k)
	}
	// range 在字符串中迭代 unicode 码点(code point)。
	// 第一个返回值是字符的起始字节位置，然后第二个是字符本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
