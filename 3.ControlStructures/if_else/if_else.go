package main

import "fmt"

//定义一个函数，比较x是否大于y，是的话返回true，否则返回false
func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}

func main() {

	//单分支结构
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	//双分支结构
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//多分支结构，else if理论来说可以无限
	//在if结构中，判断语句里可以包含初始化语句，比如给一个变量赋值，固定格式是在初始化语句后加上分号
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//调用前文定义的isGreater函数，并打印比较结果
	x,y := 1,2
	var compare = isGreater(x, y)
	fmt.Println(compare)
	fmt.Println(x,y)
}
