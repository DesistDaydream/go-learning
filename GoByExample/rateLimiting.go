// 速率限制
// 速率限制是一个重要的控制服务资源利用和质量的途径。Go通过Go协程、通道和打点器优美的支持了速率限制
package main

import (
	"fmt"
	"time"
)

func main() {
	// 假设我们想限制接收请求的处理，我们将这些请求发送给一个相同的通道
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// 这个limiter通道将每200ms接收一个值。这个是速率限制任务中的管理器
	limiter := time.Tick(time.Millisecond * 200)
	// 通过爱每次请求钱阻塞limiter通道的一个接收，我们简直自己每200ms执行一次请求
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制，我们可以通过通道缓冲来实现。
	// 这个burstyLimter通道用来进行3次临时的脉冲型速率限制。
	burstyLimiter := make(chan time.Time, 3)
	// 想将通道填充需要临时改变3次的值，做好准备
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 每200ms我们将添加一个新的值到burstyLimiter中，直到达到3个的限制
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()
	// 模拟超过5个的接入请求。它们中刚开始的3个将由于受burstyLimiter的脉冲影响
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
