// 生成进程
// 生成进程可以理解为在Go程序中运行外部命令，比如在Go程序中运行Linux命令
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	Command1()	
	Command2()
	Command3()
}


func Command1() {
	// Command函数返回值包括指定外部程序的运行路径和各种关于该程序的参数，如果想要使用参数详见Command2函数中的使用
	// Output方法，会执行调用该方法的变量中的程序，返回值为该程序的输出结果的Byte类型的切片
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
}

func Command2() {
	// 与Command1函数中想要运行的外部程序不同，这里面的外部程序还包含参数，需要显示描述程序和参数数组，而不能只是用一个程序名字符串
	// 如果想使用一个字符串生成一个关乎指定程序的完成的命令，可以使用bash命令的-c选项
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h /home")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h /home")
	fmt.Println(string(lsOut))
}

func Command3() {
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
}