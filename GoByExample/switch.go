package main

import (
	"fmt"
	"time"
)

func main() {
	// 基本的switch结构，类似于if-else结构
	i := 2
	fmt.Println("Write", i, "as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// 在一个case语句中，可以使用`,逗号`来分隔多个表达式
	// 除了case语句指定的条件外，其余全部执行default关键字后面的语句
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
	// switch关键字后面不带表达式的结构是实现if/else逻辑的另一种方式
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}
	// 类型开关(type switch)比较类型而非值。
	// 可以用来发现一个接口值的类型。
	// 在这个例子中，变量t在每个分支中会有相应的类型。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
