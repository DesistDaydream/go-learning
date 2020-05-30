// 定时器，用来在当前时间之后的某一个时刻运行Go代码
package main

import (
	"fmt"
	"time"
)

func main() {
	// NewTimer会创建一个定时器，在最少过去时间段`参数内定义的时间`后到期
	// 返回值为名为Timer的结构体指针，其中时间会被发送给一个通道C
	timer1 := time.NewTimer(time.Second * 5)
	// 直到定时器的通道C明确的发送了定时器失效的值之前，将一直阻塞
	<-timer1.C
	fmt.Println("Timer 1 expired")
	// 
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

// 第一个定时器将在程序开始后~5秒后失效，但是第二个在他没失效之前就停止了
