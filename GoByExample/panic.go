// panic意味着有些出乎意料的错误发生
// 通常用panic表示程序正常运行中不应该出现的或者我们没有处理好的错误
package main

import (
	"os"
)

func main() {
	panic("a problem")

	// 两段panic示例分开执行，否则无法看出效果，最好注释其中一个
	_, err := os.Create("/tmp/file")
	if err == nil {
		panic(err)
	}
}

// 运行程序会引起panic，输出一个错误消息和Go运行时栈信息，并且返回一个非零的状态码。
