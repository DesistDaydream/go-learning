package main

import (
	"fmt"
	"time"
)

func correct() {
	// 使用make函数创建一个新的通道。通道类型就是他们需要传递值的类型
	messages := make(chan string)

	// 使用`ChannelID <-`语法发送(send)一个新的值到通道中
	go func() { messages <- "ping" }()

	// 使用`<- ChannelID`语法从通道中接收(receives)一个值
	msg := <-messages
	fmt.Println(msg)
}

func error() {
	messages := make(chan string)

	// 如果不使用协程，则运行程序时，会报“死锁”的错,错误信息如下：
	// fatal error: all goroutines are asleep - deadlock!
	// 因为，代码是一行一行执行的，如果一个没有缓存的通道在接收数据之后，需要同步把数据发送给接收者。
	// 可是当前行的代码还没执行完，怎么能执行后面的呢，没有后面的代码，也就没有接收者，所以这就是错误产生的原因。
	// 当通道使用协程的方式运行时，就算当前时刻没有接收者，这个通过协程运行起来的通道，一会自动阻塞，并等待接收者。否则不通过协程启动通道，那么就跟普通代码一样。
	func() { messages <- "ping" }()

	// 可以使用如下方式直接输出通道内的数据，相当于fmt.Println就是通道的接收者
	fmt.Println(<-messages)
}

func buffer() {
	messages := make(chan string, 1)

	// 带缓冲的通道可以不使用协程。
	messages <- "ping"

	fmt.Println(<-messages)
}

func forChannel() {
	// 数据接受者总是面临这样的问题：何时停止等待数据？还会有更多的数据么，还是所有内容都完成了？我应该继续等待还是该做别的了？
	// 对于该问题，一个可选的方式是，持续的访问数据源并检查channel是否已经关闭，但是这并不是高效的解决方式。
	// Go 提供了 range 关键字，将其使用在 channel 上时，会自动等待 channel 的动作一直到 channel 被关闭

	// 这里使用 time 库创建一个可以持续生产数据的 Channel，每隔 1s 发送一次
	ticker := time.NewTicker(1 * time.Second)
	// 使用 for range 关键字持续从通道中消费
	for range ticker.C {
		fmt.Println("test")
		// 可以设计达到某些条件后关闭发送数据的通道。当然，也可以不关闭，这就相当于做了一个定时器，周期执行某些代码。
		// ticker.Stop()
		// break
	}

}

func main() {
	fmt.Println("1.通道正确的示例")
	// correct()
	fmt.Println("2.通道会导致死锁的示例")
	// error()
	fmt.Println("3.通道缓冲")
	// buffer()
	forChannel()

}

// 默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
// 这个特性允许我们，不使用任何其它的同步操作，来在程序结尾等待消息"ping"
