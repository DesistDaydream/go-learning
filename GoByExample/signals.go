// 信号
// 有时候我们系统Go程序能只能得处理`Linux信号`。比如我们希望当服务器收的到一个SIGTERM信号时能够自动关机，或者一个命令行工具在接收到一个SIGINT信号时停止处理输入信息
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}


/*
当我们运行这个程序时，将一直等待一个信号。使用ctrl-c，我们可以发送一个SIGINT信号，这会使得该程序打印interupt然后退出
*/