// 关闭一个通道指的是不能再想这个通道发送新值
// 这个特性可以用来给通道的接收方传达工作已经完成的信息
package main

import (
	"fmt"
)

func channelClosing() {
	// jobs通道来传递Go协程任务执行的结束信息
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 启动一个协程从`jobs通道`接收数据
	// 如果jobs已经关闭并且通道中所有的值都已经接收完毕，那么more的值为false
	// 当完成所有人物后，会给done通道发送一个通知，解除阻塞。
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 给jobs通道发送3个数据
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// 关闭jobs通道，不再接收新数据
	close(jobs)
	fmt.Println("sent all jobs")
	// 使用channelSync.go中的通道同步方法，等待任务结束
	<-done
}
