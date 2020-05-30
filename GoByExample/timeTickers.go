// 打点器，用于在固定的时间间隔重复执行代码
package main

import (
	"fmt"
	"time"
)

func main() {
	// NewTicker会创建一个打点器
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	// 将在运行1600ms后停止打点器
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
