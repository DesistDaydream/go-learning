package main

import (
	"fmt"
	"os"
	"regexp"
)

// 匹配中文
var exprCN = "[\u4e00-\u9fa5]"

func main() {
	srcFile := "test_files/test.txt"
	// 待处理文本
	fileByte, _ := os.ReadFile(srcFile)

	// 编写正则表达式
	var exprLine = `\s*\n` // 匹配空白字符任意数后接换行符，暂时搜不到 Go 中如何直接匹配空行
	// 使用正则表达式实例化一个 Regexp 对象。这里称之为编译正则表达式，也可以称之为创建正则表达式对象。
	regexp := regexp.MustCompile(exprLine)
	// 使用正则表达式对象的 ReplaceAllString 方法，将匹配到的字符串替换为换行符
	result := regexp.ReplaceAllString(string(fileByte), "\n")
	fmt.Println(result)
}
