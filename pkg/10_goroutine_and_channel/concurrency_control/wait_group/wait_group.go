package main

import (
	"fmt"
	"sync"
	// "time"
)

func sendData(ch chan string, done chan bool, wg sync.WaitGroup) {
	// 该函数把几个字符串发送给通道ch，使得这些字符串会存在一个管道中，等待输出
	ch <- "Tianjin"
	ch <- "Beijing"
	ch <- "China"

	// 当完成数据发送时，发送通知
	done <- true

	// 让 WaitGroup 计数器 -1
	wg.Done()
}

func recvData(ch chan string, done chan bool, wg sync.WaitGroup) {
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

	// 让 WaitGroup 计数器 -1
	wg.Done()
}

func main() {
	// 声明 WaitGroup 计数器
	// WaitGroup 等待一组 goroutine 完成。主 goroutine 调用 Add 来设置要等待的 goroutine 的数量。
	// 然后每个 goroutine 运行并在完成时调用 Done。 同时，Wait 可用于阻塞，直到所有 goroutine 完成。
	var wg sync.WaitGroup
	// 为 WaitGroup 计数器 +2
	wg.Add(2)

	// ch 通道用来传递数据；done 通道用来传递任务是否完成的消息。
	ch := make(chan string)
	done := make(chan bool)

	// 把通道ch作为参数传递给两个协程函数
	// 可以理解为把通道的两端分别连接到两个协程函数上
	go sendData(ch, done, wg)
	go recvData(ch, done, wg)

	// 如果不让 main() 等待，则无任何输出，或者报错 panic: send on closed channel。因为协程是并发运行，不用等代码运行完成就会执行后续代码。
	// 如果后续代码执行完了，协程中的代码还没执行完成，就会没有任何输出。
	// 如果后续代码中包含了关闭通道的操作，那么程序将会 panic，并报错 send on closed channel
	// 可以使用time等待一秒
	// time.Sleep(time.Second)
	//
	// 通常，我们使用 sync.WaitGroup 来让程序等待协程完成，就像前面几行代码一样，上面使用 time.Sleep() 仅仅就是进行一下说明。
	// 平时我们会将 wg.Wait() 写在 var wg sync.WaitGroup 的下一行，并在语句前加上 defer。即：
	// ```
	// var wg sync.WaitGroup
	// defer wg.Wait()
	// ```
	wg.Wait()
}
