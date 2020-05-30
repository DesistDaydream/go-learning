// 结构体是带类型的字段(fields)集合。这在组着数据时非常有用
package main

import (
	"fmt"
)

// person结构体包含了`name`和`age`两个字段
type person struct {
	name string
	age  int
}

func main() {
	// 使用该语法创建新的结构体元素
	fmt.Println(person{"Bob", 20})
	// 可以在初始化一个结构体元素时指定字段名
	fmt.Println(person{name: "Alice", age: 30})
	// 省略的字段将被初始化为`零`值
	fmt.Println(person{name: "Fred"})
	// `&`前缀生成一个结构体指针
	fmt.Println(&person{name: "Ann", age: 40})
	// 使用`.`符号来访问结构体内的某个字段
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	// 可以对结构体指针使用`.`，指针会被自动解引用
	sp := &s
	fmt.Println(sp.age)
	// 结构体是可变的
	sp.age = 51
	fmt.Println(sp.age)
}
