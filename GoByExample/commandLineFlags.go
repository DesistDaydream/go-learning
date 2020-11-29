// 命令行标志
// 命令行标志是命令行程序制定选项的常用方式。比如`ls -l`命令中`-l`就是ls程序的命令行标志
package main

import (
	"flag"
	"fmt"
)

func main() {
	var svar string
	// 基本的flag申明仅支持string、int、bool类型。这里声明一个名为word的标志、默认值为foo、并带有一个简短的描述
	// 函数会返回一个字符串指针，该指针指向flag所对应的值，初始值为定义的默认值
	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	// 与上面的不同，该函数无返回值，flag的值会保存在函数的第一个参数的指针中
	flag.StringVar(&svar, "svar", "bar", "a string var")
	// 所有flag声明完成以后，调用`flag.Parse()`来解析命令行，否则命令行标志不起作用，只能自己使用默认值
	flag.Parse()

	// 我们将仅输出解析的选项以及后面的位置参数
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

/*
可以使用如下方式测试该程序的运行结果
使用 -h 或者 --help 标志来得到这个程序的帮助文本
1. ./commandLineFlags -h
注意：单横线 - 也可以用双横线 --
2. ./commandLineFlags -word opt
3. ./commandLineFlags --word opt a1
4. ./commandLineFlags -word=opt a1 a2 a3
5. ./commandLineFlags -word=opt a1 a2 a3 -numb=7
如果提供一个没有定义的flag，程序会输出一个错误信息，并再次显示帮助文本
6. ./commandLineFlags -wat
*/
