// 执行进程
// 在spawningProcesses.go中，在Go程序中访问外部程序。但是有时候我们只想用其他的程序来完全替代当前的GO程序。这时候就可以用该篇代码来实现
package main

import (
	"os"
	"os/exec"
	"syscall"
)

func execingProcesses() {
	// 取得将要执行的程序的绝对路径
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// 为该程序提供一些参数
	args := []string{"ls", "-a", "-l", "-h"}

	// 为程序提供环境变量
	env := os.Environ()

	// 执行命令，包括路径，程序与参数，环境变量
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

/*
该Go程序的执行结果与`ls -a -l -h`命令的执行结果相同
*/
