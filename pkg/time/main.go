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

func timeFormat() {
	currentTime := time.Now()
	// 在 time 包的 format.go 文件中，有很多官方提供的格式，这些格式都是字符串常量，我们可以直接调用。
	fmt.Println("ANSIC : ", currentTime.Format(time.ANSIC))
	// 除了官方提供的格式，我们还可以自动定义时间格式。
	// 自己定义时间格式时，只需要以 2006 年 1 月 2 日 15 点 04 分 05 秒为基准即可，这是 Go 语言诞生的时间
	fmt.Println("当前时间  : ", currentTime)
	fmt.Println("当前时间字符串: ", currentTime.String())
	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))
	fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000"))
	fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000000"))
	fmt.Println("ShortNum Month : ", currentTime.Format("2006-1-02"))
	fmt.Println("LongMonth : ", currentTime.Format("2006-January-02"))
	fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))
	fmt.Println("ShortYear : ", currentTime.Format("06-Jan-02"))
	fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))
	fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))
	fmt.Println("ShortDay : ", currentTime.Format("Mon 2006-01-2"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 pm"))

	fmt.Println("当前时间的前一天", currentTime.AddDate(0, 0, -1).Format("2006-01-02 15:04:05"))
}

func main() {
	// timer()
	// ticker()
	// timestamp()
	timeFormat()
}
