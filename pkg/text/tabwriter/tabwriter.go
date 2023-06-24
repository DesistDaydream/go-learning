package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/olekukonko/tablewriter"
)

// tabwriter 包可以将输出以表格式的对其方式展示出来
func tabwriterDemo() {
	// 创建一个新的tabwriter.Writer，并指定输出到标准输出
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', tabwriter.Debug)

	// 写入数据，使用制表符分隔列
	// TODO:
	fmt.Fprintln(w, "min\tAge\tCity")
	fmt.Fprintln(w, "John Doe\t30\tNew York")
	fmt.Fprintln(w, "Jane Smith\t25\tSan Francisco")
	fmt.Fprintln(w, "Bob Johnson\t35\tChicago")

	// 刷新缓冲区，确保所有数据都写入
	// 注意：确保所有数据都写入到 tabwriter.Writer 后，在执行 Flush() 方法，否则无法让每列对齐。
	w.Flush()
}

func tablewriterDemo() {
	data := [][]string{
		{"张三", "25", "程序员"},
		{"李四", "30", "教师"},
		{"王五", "35", "医生"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"姓名", "年龄", "职业"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
}

func main() {
	// text/tabwriter 包示例
	tabwriterDemo()

	// 注意：若数据变成中文，则表格将会错位，除非只有最后一列是中文，
	// 这是因为 text/tabwriter 库是按照 UTF-8 编码的文本单元来处理输入的数据的，而中文字符在 UTF-8 编码中占用的字节数不一定相同，所以对齐的时候可能会出现偏差。
	// 如果要解决这个问题，可以使用第三方库，如 github.com/olekukonko/tablewriter，它可以支持中文字符的对齐。
	tablewriterDemo()
}
