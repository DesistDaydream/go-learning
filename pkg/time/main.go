package main

import (
	"fmt"
	"time"
)

func timer() {
	// NewTimer会创建一个定时器，在最少过去时间段`参数内定义的时间`后到期
	// 返回值为名为Timer的结构体指针，其中时间会被发送给一个通道C
	timer1 := time.NewTimer(time.Second * 5)
	// 直到定时器的通道C明确的发送了定时器失效的值之前，将一直阻塞
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 由于以协程方式启动，还没等到失效，就运行了 Stop()
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// 第一个定时器将在程序开始后~5秒后失效，但是第二个在他没失效之前就停止了
}

func ticker() {
	var i int
	// NewTicker 创建一个打点器
	// Ticker 常用来与 for 循环结合使用,在一个死循环中，每间隔一定时间，就执行一遍代码
	ticker := time.NewTicker(time.Second * 3)

	for {
		switch i {
		// 用于退出 for{} 循环
		case 5:
			return
		// 每间隔3秒，输出当前时间
		default:
			<-ticker.C
			fmt.Println(time.Now())
			// 当输出5次后，即退出 for{} 循环
			i = i + 1
		}
	}
}

func timestamp() {
	fmt.Printf("时间戳（秒）：%v;\n", time.Now().Unix())
	fmt.Printf("时间戳（纳秒）：%v;\n", time.Now().UnixNano())
	fmt.Printf("时间戳（毫秒）：%v;\n", time.Now().UnixNano()/1e6)
	fmt.Printf("时间戳（纳秒转换为秒）：%v;\n", time.Now().UnixNano()/1e9)
}

func main() {
	// timer()
	// ticker()
	timestamp()
}
