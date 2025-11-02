package main

import (
	"fmt"
	"time"
)

func channelCorrect() {
	// 使用 make 函数创建一个新的通道。通道类型就是他们需要传递值的类型
	messages := make(chan string)

	// 使用 `ChannelID <-` 语法发送(send)一个新的值到通道中
	go func() { messages <- "ping" }()

	// 使用 `<- ChannelID` 语法从通道中接收(receives)一个值
	msg := <-messages
	fmt.Println(msg)
}

func channelError() {
	messages := make(chan string)

	// 如果不使用协程，则运行程序时，会报“死锁”的错,错误信息如下：
	// fatal error: all goroutines are asleep - deadlock!
	// 因为，代码是一行一行执行的，如果一个没有缓存的通道在接收数据之后，需要同步把数据发送给接收者。
	// 可是当前行的代码还没执行完，怎么能执行后面的呢，没有后面的代码，也就没有接收者，所以这就是错误产生的原因。
	// 当通道使用协程的方式运行时，就算当前时刻没有接收者，这个通过协程运行起来的通道，会自动阻塞，并等待接收者。否则不通过协程启动通道，那么就跟普通代码一样。
	func() { messages <- "ping" }()

	// 可以使用如下方式直接输出通道内的数据，相当于fmt.Println就是通道的接收者
	fmt.Println(<-messages)
}

func channelBuffer() {
	messages := make(chan string, 1)

	// 带缓冲的通道可以不使用协程。
	messages <- "ping"

	fmt.Println(<-messages)
}

// 检测通道是否还有数据的语法
func isOkChannel() {
	messages := make(chan string, 1)

	messages <- "ping"

	// 可以从通道接收两个参数，第一个是值，第二个是是否接收到值
	msg, ok := <-messages
	if ok {
		fmt.Println(msg)
	}
}

func rangeChannel() {
	// 数据接收者总是面临这样的问题：何时停止等待数据？还会有更多的数据么，还是所有内容都完成了？我应该继续等待还是该做别的了？
	// 对于该问题，一个可选的方式是，持续的访问数据源并检查channel是否已经关闭，但是这并不是高效的解决方式。
	// Go 提供了 range 关键字，将其使用在 channel 上时，会自动等待 channel 的动作一直到 channel 被关闭

	// 这里使用 time 库创建一个可以持续生产数据的 Channel，每隔 2s 发送一次
	ticker := time.NewTicker(2 * time.Second)
	// 使用 for range 关键字持续从通道中消费数据
	for channelData := range ticker.C {
		fmt.Println("已消费 Channel 中的数据: ", channelData)

		// 可以设计达到某些条件后关闭发送数据的通道。当然，也可以不关闭，这就相当于做了一个定时器，周期执行某些代码。
		// ticker.Stop()
		// break
	}

	// 也可以省略前面用于存储通道数据的变量，写成这样
	// for range ticker.C {
	// 	fmt.Println("消费 Channel 中的数据")
	// }

}

func ChannelAndSwitch() {
	doneChannel := make(chan struct{})
	timeChannel := time.NewTicker(1 * time.Second).C

	// 如果两个 channel 都没有数据，select 会阻塞等待
	// 如果只有一个 channel 有数据，执行对应的 case
	// 如果两个 channel 同时都有数据，Go 会随机选择一个 case 执行
	// 执行完一个 case 后，select 语句结束（不会执行另一个 case）
	go func() {
		for {
			select {
			// 当 doneChannel 通道有数据可读或被关闭时，执行该分支
			case <-doneChannel:
				return
			// 当 timeChannel 通道有数据可读时，执行该分支
			case timeData := <-timeChannel:
				fmt.Println("已消费 Channel 中的数据: ", timeData)
			}
		}
	}()
	// 这种模式常用于"带取消功能的数据处理"。在其他函数逻辑中关闭 doneChannel 即可接触 goroutine 中的 for 循环，从而结束 goroutine。

	time.Sleep(5 * time.Second)

	// 关闭 doneChannel 通道，触发 goroutine 中的 case <-doneChannel: 分支
	close(doneChannel)
}

func main() {
	fmt.Println("通道会导致死锁的错误示例")
	channelError()
}

// 默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕。
// 这个特性允许我们，不使用任何其它的同步操作，来在程序结尾等待消息"ping"
