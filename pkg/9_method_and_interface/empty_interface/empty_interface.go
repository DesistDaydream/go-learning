package main

import "fmt"

var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

// 空接口
// 空接口是没有任何方法的接口，任何类型都实现了空接口，所以空接口可以表示任何类型
// 空接口的定义为：interface{}
// 空接口的用途：
// 1. 作为函数的参数，表示该函数可以接受任何类型的参数
// 2. 作为函数的返回值，表示该函数可以返回任何类型的值
// 3. 作为变量，表示该变量可以接受任何类型的值
func main() {
	var val Any
	// 空接口也可以写成下面这样
	// var val interface{}

	val = "5"
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)

	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
