// Go协程(goroutine)在执行上来说是轻量级的线程。
package main

import (
	"fmt"
)

func fa(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goroutines() {
	// 一般情况下，正常调用函数。同步地(synchronously)运行
	fa("direct")
	// 在一个Go协程中调用这个函数。这个新的Go协程将会并发(concurrently)执行这个函数
	go fa("goroutine")
	// 也可以为匿名函数启动一个Go协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// 上面两个Go协程在独立的Go协程中异步(asynchronously)运行。
	// 程序直接运行到这一行，这里需要在程序退出前按下任意键结束
	fmt.Scanln()
	fmt.Println("done")
}
