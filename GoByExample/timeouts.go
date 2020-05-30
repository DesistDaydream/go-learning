package main

import (
	"fmt"
	"time"
)

func main() {
	// 假设执行一个外部调用，并在2秒后通过通道c1返回执行结果
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	// 实现超时操作，由于select默认处理第一个已经准备好的接收操作
	// 如果这个操作超过了允许的1秒，将会执行超时case
	select {
	// 等待结果
	case res := <-c1:
		fmt.Println(res)
	// 等待超时时间1秒后发送的值
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// 超时时间3秒，将会成功从c2接收到并打印
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
