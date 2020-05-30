// 命令行参数
// 命令行参数是指定程序运行参数的一个常见方式。比如go run test.go,程序go使用了run和test.go两个参数
package main

import (
	"fmt"
	"os"
)

func main() {
	// `os.Args`变量提供原始命令行参数访问功能。
	// os.Args变量是一个字符串切片，保存着程序运行的命令行参数，且该切片中的第一个元素是该程序的绝对路径
	// 并且`os.Args[1:]`保存该程序的所有参数
	// argsWithProg为文件名、命令行参数1、...、命令行参数N
	argsWithProg := os.Args
	// 输出程序的所有参数
	argsWithoutProg := os.Args[1:]
	// 直接使用索引来获得对应位置的参数的值,输出第四个参数，如果有该行代码，则程序运行的时候至少要额外输入3个参数
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
