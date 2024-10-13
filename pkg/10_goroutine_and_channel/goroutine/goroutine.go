package main

import (
	"fmt"
	"math"
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

func timeSleepGorourine() {
	// 调用了 go hello() 之后，程序控制没有等待 hello 函数结束，
	// 立即返回到了代码下一行，打印 main function。
	// 接着由于没有其他可执行的代码，Go 主协程终止，
	// 于是 hello 函数就没有机会运行了。
	go hello("timesleep")

	// 为了避免避免程序无法运行 hello()，需要让程序等待一段时间。
	time.Sleep(1 * time.Second)
	// 也可以使用下面的方式让程序在手动按回车才结束。
	// fmt.Scanln()
	fmt.Println("main function")
}

func waitGroupGoroutine() {
	// 为 WaitGroup 计数器 +1
	wg.Add(1)

	go hello("waitgroup")

	// 等待 WaitGroup 计数器归零，归零后，wg.Wait() 释放，并继续执行后面的代码、
	wg.Wait()

	fmt.Println("main function")
}

// 该操作执行一段时间后将会报错：
// panic: too many concurrent operations on a single file or socket (max 1048575)
func maxGoroutine() {
	var wg sync.WaitGroup
	for i := 0; i < math.MaxInt32; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("并发数量：%d\n", i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}

// Go 协程的最基本示例，demo() 函数在运行时，其实并不会输出内容
// 因为 Go 协程的效果是：使用 go 关键字调用 hello() 之后，程序并不会等待 hello 函数结束，
// 而是立即开始执行后面的代码，接着由于没有其他可执行的代码，Go 主协程终止，
// 于是有可能出现 hello 的逻辑还没有处理完，程序本身就自动退出了
// 所以，使用在使用 Go 协程时，往往不能独立使用 go 关键字，而是要搭配其他机制，比如 WaitGroup、etc.
func demo() {
	go hello("demo")
	fmt.Println("这里应该看不到 demo 的输出")
}

func main() {
	// 一个独立的 Go 协程基本演示。Note: 这函数是错误的用法，仅仅展示协程的基本概念
	demo()
	// 通过睡眠，让 main() 等待协程完成
	timeSleepGorourine()
	// 通过 WaitGroup，让 main() 等待协程完成
	waitGroupGoroutine()
	// 最大协程
	maxGoroutine()
}
