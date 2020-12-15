package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "hello world"
	// 判断变量s是否包含字符hello
	fmt.Println(strings.Contains(s, "hello"))
	// 输出o在变量s的索引号
	fmt.Println(strings.Index(s, "o"))

	ss := "1#2#345"
	// 切割字符串
	splitedStr := strings.Split(ss, "#")
	fmt.Println(splitedStr)
	// 合并字符串
	fmt.Println(strings.Join(splitedStr, "*"))
	// 判断s是否以he开头,判断s是否以ld结尾
	fmt.Println(strings.HasPrefix(s, "he"), strings.HasSuffix(s, "ld"))
	// 将int类型转换为string类型
	fmt.Println(strconv.Itoa(20))
	// 将string类型转换为int类型
	fmt.Println(strconv.Atoi("10"))
}
