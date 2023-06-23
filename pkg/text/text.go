package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// tabwriter 包可以将输出以表格式的对其方式展示出来
func tabwriterDemo() {
	// 创建一个新的tabwriter.Writer，并指定输出到标准输出
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.Debug)

	// 写入数据，使用制表符分隔列
	// TODO: 若数据变成中文，则表格将会错位，这个问题如何解决？
	fmt.Fprintln(w, "min\tAge\tCity")
	fmt.Fprintln(w, "John Doe\t30\tNew York")
	fmt.Fprintln(w, "Jane Smith\t25\tSan Francisco")
	fmt.Fprintln(w, "Bob Johnson\t35\tChicago")

	// 刷新缓冲区，确保所有数据都写入
	w.Flush()
}

func main() {
	tabwriterDemo()
}
