// 通道同步。可以使用通道来同步Go协程间的执行状态
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	// done通道将被用于通知其他Go协程这个函数已经工作完毕
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// 发送一个值来通知我们执行完成
	done <- true
}
func main() {
	done := make(chan bool, 1)
	// 运行一个Go协程，并给与用于通知的通道
	go worker(done)
	// 程序将在done通道接收到数据之前，一直阻塞而不让主程序结束
	// 如果把该行代码删除，程序有可能会在worker函数还没开始运行时就结束了
	<-done
}
