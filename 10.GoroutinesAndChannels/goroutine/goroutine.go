package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(m string) {
	if m == "waitgroup" {
		// 让 WaitGroup 计数器 -1
		defer wg.Done()
	}

	i := 0
	for {
		fmt.Printf("Hello world goroutine,%v\n", i)
		i += 1
		if i == 3 {
			return
		}
	}
}

func timeSleepGotourine() {
	// 调用了 go hello() 之后，程序控制没有等待 hello 协程结束，
	// 立即返回到了代码下一行，打印 main function。
	// 接着由于没有其他可执行的代码，Go 主协程终止，
	// 于是 hello 协程就没有机会运行了。
	go hello("timesleep")

	// 调用了 time 包里的函数 Sleep，该函数会休眠执行它的 Go 协程。
	// 在这里，我们使 Go 主协程休眠了 1 秒。因此在主协程终止之前，
	// 调用 go hello() 就有足够的时间来执行了。
	// 该程序首先打印 Hello world goroutine，等待 1 秒钟之后，接着打印 main function。
	time.Sleep(1 * time.Second)
	// 也可以使用下面的方式让程序在手动按回车才结束。
	// fmt.Scanln()
	fmt.Println("main function")
}

func waitGroupGotoutine() {
	// 为 WaitGroup 计数器 +1
	wg.Add(1)

	go hello("waitgroup")

	// 等待 WaitGroup 计数器归零，归零后，wg.Wait() 释放，并继续执行后面的代码、
	wg.Wait()

	fmt.Println("main function")
}

func main() {
	// 通过睡眠，让 main() 等待协程完成
	timeSleepGotourine()
	// 通过 WaitGroup，让 main() 等待协程完成
	waitGroupGotoutine()
}
