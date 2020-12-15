// 通道选择器可以同时等待多个通道操作。
// Go协程和通道以及选择器的结合是Go的一个强大特性。
package main

import (
	"fmt"
	"time"
)

func channelSelect() {
	// 后续代码会从这两个通道中选择
	c1 := make(chan string)
	c2 := make(chan string)
	// 各个通道将在若干时间后接收一个值。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	// 使用select关键字来同时等待这两个值，并打印各自接收到的值
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
