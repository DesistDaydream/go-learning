package main

import (
	"fmt"
	"time"
)

type Human struct {
	Name string
}

var people = Human{Name: "zhangsan"}

func main() {
	// 打印结构体时，会添加字段名
	fmt.Printf("%+v\n", people)
	// Go 语法表示
	fmt.Printf("%#v\n", people)

	for i := 0; i < 10; i++ {
		fmt.Printf("\033[2K \r 第 %v 页 \033[0m", i)
		time.Sleep(100 * time.Millisecond)
	}

	// 0x1B 是 16 进制表示、\033 是 8 进制表示
	fmt.Printf("%c[33;1m 黄色字体加粗 %c[0m\n", 0x1B, 0x1B)
	fmt.Printf("%c[33;4m 黄色字体下划线 %c[0m\n", 0x1B, 0x1B)
	fmt.Printf("%c[42;30;1m 绿色背景黑色字体加粗 %c[0m\n", 0x1B, 0x1B)
}
