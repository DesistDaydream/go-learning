package main

import (
	"fmt"
	// "time"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	// 如果下面一行代码，编译会报错，因为信道没有接收到任何数据，后续代码被阻塞不会执行
	done <- true
}
func channel() {
	// 创建了一个bool类型的信道，名为done的变量
	done := make(chan bool)
	// 把done作为参数传递给hello协程，使得hello()函数可以在函数体内部使用done通道
	go hello(done)
	// 通过信道 done 接收数据。这一行代码发生了阻塞，
	// 除非有协程向 done 写入数据，否则程序不会跳到下一行代码。
	<-done
	fmt.Println("main function")

}
