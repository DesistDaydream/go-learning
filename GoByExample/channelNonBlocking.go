package main

import (
	"fmt"
)

func channelNonBlocking() {
	messages := make(chan string)
	signals := make(chan bool)
	// 这是一个非阻塞接收的例子，如果在messages中存在
	// 然后select将这个值带入<-messages case中；如果不是
	// 就直接到default分支中
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	// 一个非阻塞发送的实现方法和上面一样
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	// 可以在default前使用多个case子句来实现一个多路的非阻塞的选择器
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
