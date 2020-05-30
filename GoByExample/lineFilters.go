/*
行过滤器
行过滤器用于读取标准输入流的输入，处理该输入，然后将得到的一些结果输出到标准输出
grep和sed命令时常见的行过滤器
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 对标准输入使用一个带缓冲的扫描器，以便使用Scan方法来直接读取一行
	scanner := bufio.NewScanner(os.Stdin)

	// 每次调用该方法，即可读取scanner的下一行
	for scanner.Scan() {
		// ToUpper会把获取的到的内容转换成大写字母
		// Text返回当前的token，现在是输入的下一行
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}
