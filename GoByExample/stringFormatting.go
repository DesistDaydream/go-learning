// 字符串格式化
package main

import (
	"fmt"
	"os"
)

type point struct {
	s, y int
}

func main() {
	p := point{1, 2}
	// 使用`%v`来自动判断变量的类型
	fmt.Printf("%v\n", p)
	// 如果值是一个结构体，则`%+v`格式输出的内容将包括结构体的字段名
	fmt.Printf("%+v\n", p)
	// 该格式输出运行源码的片段
	fmt.Printf("%#v\n", p)
	// 输出将要打印的值的类型
	fmt.Printf("%T\n", p)
	// 格式输出布尔值
	fmt.Printf("%t\n", true)
	// 十进制格式
	fmt.Printf("%d\n", 123)
	// 二进制格式
	fmt.Printf("%b\n", 14)
	// 指定证书对应的字符
	fmt.Printf("%c\n", 33)
	// 十六进制
	fmt.Printf("%x\n", 456)
	// 浮点数
	fmt.Printf("%f\n", 78.9)
	// 科学计数法
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// 基本字符串的输出
	fmt.Printf("%s\n", "\"string\"")
	// 不进行转义直接输出
	fmt.Printf("%q\n", "\"string\"")
	// 输出base-16编码的字符串，每个字节使用2个字符表示
	fmt.Printf("%x\n", "hex this")
	// 输出一个指针的值
	fmt.Printf("%p\n", &p)

	// 当输出数字的时候，经常想要控制输出结果的宽度和精度，可以使用%后面使用数字来控制输出宽度
	// 默认结果使用右对齐且通过空格来填充空白部分
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// 可是指定浮点型的输出宽度，同时也可以通过`宽度.精度`语法来指定输出的精度
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// 要左对齐，使用`-`表示
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// 控制字符串输出时的宽度
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// S开头的Printf会格式化并返回一个字符串而不带任何输出
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// F开头的Printf(format+print)会格式化输出到指定的地方，比如这里是输出到`os.Stderr`
	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
