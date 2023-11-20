package main

import (
	"fmt"
	"time"
)

func main() {
	switch_1()
	switch_2()
}

func switch_1() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

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

func switch_2() {
	i := 3
	//给定一个变量，当switch中分支的变量i满足某个条件时则执行代码块
	switch i {
	//当i等于0时，执行完case 0的代码块不退出switch结构，执行下面一个分支的代码块，
	case 0:
		fallthrough
	//当i等于1时，打印变量i的值
	case 1:
		fmt.Println(i)
	case 2:
		fmt.Println(i)
	//当i的条件不满足上述所有时，执行default分支的代码块
	default:
		fmt.Println(i)
	}

	//可以在switch进行过判断，当任一分支的条件测试结果为真，则执行代码块
	num1 := 5
	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}

	//可以在switch结构开头进行初始化，下面是对a与b变量进行初始化
	var t int
	switch a, b := 1, 2; {
	case a < b:
		t = -1
	case a == b:
		t = 0
	case a > b:
		t = 1
	}
	fmt.Println(t)
}
