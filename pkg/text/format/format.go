package main

import (
	"fmt"
	"strconv"
	"time"
)

type Human struct {
	Name string
}

var people = Human{Name: "zhangsan"}

// ANSI 演示
func AnsiDemo() {
	for i := 0; i < 10; i++ {
		// [K 用于清除光标到行尾的内容；\r 用于将光标移动到行首。通过这种控制器，让终端可以始终只显示一行。
		fmt.Printf("\033[  2  K \r 第 %v 页 \033[0m", i)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\n") // 由于上面的终端控制最后不知道怎么换行，会跟后面的终端输出到一行，所以手动换行

	// 颜色控制
	// 0x1B 与 \033 的作用一样，用来表示 ACII 的 ESC 字符。只不过 0x1B 是 16 进制形式、\033 是 8 进制形式
	fmt.Printf("%c[33;1m 黄色字体加粗 %c[0m\n", 0x1b, 0x1b)
	fmt.Printf("%c[33;4m 黄色字体下划线 %c[0m\n", 0x1b, 0x1b)
	fmt.Printf("%c[42;30;1m 绿色背景黑色字体加粗 %c[0m\n", 0x1b, 0x1b)
}

func colorDemo() {
	flag := 0
	str := "KURISU "
	for r := 255; r >= 0; r -= 2 {
		for g, b := 0, 255; g < 255 && b >= 1; g += 1 {
			if flag >= len(str) {
				flag = 0
			}
			a := str[flag]
			flag++
			fmt.Printf("\x1b[48;2;%s;%s;%sm\x1b[38m%s\x1b[0m", strconv.Itoa(r), strconv.Itoa(g), strconv.Itoa(b), string(a))
			b -= 1
		}
		fmt.Printf("\x1bE")
	}
}

func main() {
	// 打印结构体时，会添加字段名
	fmt.Printf("%+v\n", people)
	// Go 语法表示
	fmt.Printf("%#v\n", people)

	AnsiDemo()

	// 清屏的多种实现方式，类似 clear 命令或 ctrl + l 快捷键
	// fmt.Printf(" \x1b[2J")
	// fmt.Printf(" \x1bc")

	colorDemo()
}
