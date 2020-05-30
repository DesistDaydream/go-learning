package main

import (
	"fmt"
	"sync"
	// "time"
)

var (
	wg   sync.WaitGroup
	done = make(chan bool)
)

func sendData(ch chan string) {
	// 该函数把几个字符串发送给通道ch，使得这些字符串会存在一个管道中，等待输出
	ch <- "Tianjin"
	ch <- "Beijing"
	ch <- "China"
	// 当完成数据发送时，通知
	done <- true

	wg.Done()
}

func getData(ch chan string) {
I:
	for {
		select {
		// 把通道中的字符串都赋值给变量input
		case input := <-ch:
			fmt.Println(input)
		// 当通道通知完成数据传输时，跳出循环
		case <-done:
			break I
		}
	}

	wg.Done()
}

func main() {
	wg.Add(2)

	ch := make(chan string)
	// 把通道ch作为参数传递给两个协程函数
	// 可以理解为把通道的两端分别连接到两个协程函数上
	go sendData(ch)
	go getData(ch)
	// 如果不让main()等待，则无任何输出.可以使用time等待一秒，也可以使用WaitGroup让效果更好
	// time.Sleep(time.Second)
	wg.Wait()
}
