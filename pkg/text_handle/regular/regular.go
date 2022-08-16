package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	srcFile := "test_files/test.txt"
	// 待处理文本
	fileByte, _ := os.ReadFile(srcFile)

	// 编写正则表达式
	expr := `^(\s*)\n` // 匹配空行，这个表达式有问题，待验证，暂时搜不到 Go 中如何删除空白行的方式
	// 使用正则表达式实例化一个 Regexp 对象。这里称之为编译正则表达式，也可以称之为创建正则表达式对象。
	regexp := regexp.MustCompile(expr)
	// 使用正则表达式对象的 ReplaceAllString 方法，将匹配到的字符串替换为空
	result := regexp.ReplaceAllString(string(fileByte), "")
	fmt.Println(result)
}
