// 原子计数器
// Go中最主要的状态管理方式使通过通道间的沟通来完成
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func atomicCounters() {
	// 使用一个无符号整数来表示这个计数器
	var ops uint64 = 0
	// 启动50个Go协程，对计数器每隔1ms进行一次加一操作
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用AddUint64来让计数器自动增加，使用&语法来给出ops的内存地址
				atomic.AddUint64(&ops, 1)
				// 允许其它Go协程的执行
				runtime.Gosched()
			}
		}()
	}
	// 等待1秒让ops的自加操作执行一会
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
