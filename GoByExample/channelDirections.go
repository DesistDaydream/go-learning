// 可以让通知只用来发送或接收值
package main

import (
	"fmt"
)

// 一个只允许发送数据的通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 通道pings来接收数据，通道pongs来发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func channelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
