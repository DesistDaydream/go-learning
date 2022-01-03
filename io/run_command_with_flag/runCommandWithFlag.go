package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"flag"
)

func main() {
	var fileName string
	// 使用flag包的函数来引入命令行给定的参数到程序中，使用格式为：`go run runCommand -f 赋值给fileName变量的值`,没有默认文件名
	flag.StringVar(&fileName, "f", " ", "请指定文件名")
	// 解析flag包中的函数的各种已定义的标志，如果不解析则无法使用
	flag.Parse()

	switch os.Args[1] {
	case "-f":
		ReadFile(fileName)
	default:
		fmt.Println("请指定一个选项")
	}
}

func ReadFile(fileName string) {	
	inputSlice := make([]string, 0, 100)
	inputFile, inputError := os.Open(fileName)
	if inputError != nil {
		fmt.Println("no this file")
		return
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)

	for {
		// 从读取器中按行读取内容，不包含换行符
		inputByte, _, readerError := inputReader.ReadLine()
		// 注意判断错误值的位置，如果放在循环最后，则ReadLine还会输出一个空行，这样会出现问题
		if readerError == io.EOF {
			break
		}
		// 把从读取器中获取的byte类型的数据转换成string类型以供人类阅读。再使用Split把字符串转换为切片。
		inputString := string(inputByte)
		inputSlice = strings.Split(inputString, " ")
		RunCommand(inputSlice)
	}
}
func RunCommand(a []string) {
	// 为了在命令参数中引用变量，所以需要先把命令参数赋值给变量，Command函数中再引用该参数的变量
	param := fmt.Sprintf("touch %s && echo '%s' > %s", a[0], a[1], a[0])
	cmd := exec.Command("/bin/bash", "-c", param)
	// 把命令定义好后，使用Output方法来执行命令,通过判断返回值的错误信息来判断命令是否执行成功。
	// 也可是直接使用cmd.Output()进行命令执行，只不过这样无法获取函数返回值，也就无法判断命令是否执行成功。
	if _, err := cmd.Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(a)
	cmd.Run()
}
