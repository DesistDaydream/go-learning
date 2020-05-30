package main

import (
	"fmt"
	"sync"
	// "time"
)

var (
	wg   sync.WaitGroup	
)

func sendData(ch chan string, done chan bool) {
	// 该函数把几个字符串发送给通道ch，使得这些字符串会存在一个管道中，等待输出
	ch <- "Tianjin"
	ch <- "Beijing"
	ch <- "China"

	// 当完成数据发送时，发送通知
	done <- true

	wg.Done()
}

func getData(ch chan string, done chan bool) {
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
	// ch通道用来传递数据；done通道用来传递任务是否完成的消息。
	ch := make(chan string) 
	done := make(chan bool)

	// 把通道ch作为参数传递给两个协程函数
	// 可以理解为把通道的两端分别连接到两个协程函数上
	go sendData(ch,done)
	go getData(ch,done)
	// 如果不让main()等待，则无任何输出.因为协程是并发运行，不用等代码运行完成就会执行后续代码。如果main里的代码都执行完了，协程中的代码还没执行，就会没有任何输出。
	// 可以使用time等待一秒，也可以使用WaitGroup让效果更好
	// time.Sleep(time.Second)
	wg.Wait()
}
