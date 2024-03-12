package main

import (
	"fmt"
	"os"
	"os/exec"
)

// 运行命令演示
func runCommandDemo() {
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("命令执行错误%v", err)
		return
	}
	fmt.Printf("%s", out)
}

// 使用 os.StartProcess 函数运行外部命令
func runCommandWithStartProcess() {
	//Environ()会返回当前环境的环境变量，格式为键值对切片
	env := os.Environ()
	// ProcAttr 结构体里保存被 StartProcess()`函数用于一个新进程的属性
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// StartProcess()函数使用提供的属性、程序名、命令行参数开始一个新进程
	// 该函数是一个低水平的接口，os/exec包的函数提供了高水平接口，应尽量使用该包
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("错误的命令%v启动进程!", err) //
		os.Exit(1)
	}
	fmt.Printf("这个进程的ID是：%v", pid)
}

// 使用 exec.Command 函数运行外部命令
func runCommandWithCommand() {
	// exec.Command()函数返回一个类型为*exec.Cmd的结构体指针
	// 该指针用于使用给出的参数执行第一个参数指定的程序
	cmd := exec.Command("mkdir", "-p", "1/2/3")
	// 让Run()方法作用在上面变量上，执行变量中的命令，并阻塞直到完成
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)
}
