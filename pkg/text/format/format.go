package main

import (
	"fmt"
	"time"
)

type Human struct {
	Name string
}

var people = Human{Name: "zhangsan"}

// ANSI 演示
func AnsiDemo() {
	for i := 0; i < 10; i++ {
		// [K 用于清除光标到行位的内容；\r 用于将光标移动到行首。通过这种控制器，让终端可以始终只显示一行。
		fmt.Printf("\033[2K \r 第 %v 页 \033[0m", i)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\n") // 由于上面的终端控制最后不知道怎么换行，会跟后面的终端输出到一行，所以手动换行

	// 颜色控制
	// 0x1B 与 \033 的作用一样，用来表示 ACII 的 ESC 字符。只不过 0x1B 是 16 进制形式、\033 是 8 进制形式
	fmt.Printf("%c[33;1m 黄色字体加粗 %c[0m\n", 0x1B, 0x1B)
	fmt.Printf("%c[33;4m 黄色字体下划线 %c[0m\n", 0x1B, 0x1B)
	fmt.Printf("%c[42;30;1m 绿色背景黑色字体加粗 %c[0m\n", 0x1B, 0x1B)
}

func main() {
	// 打印结构体时，会添加字段名
	fmt.Printf("%+v\n", people)
	// Go 语法表示
	fmt.Printf("%#v\n", people)

	AnsiDemo()
}
