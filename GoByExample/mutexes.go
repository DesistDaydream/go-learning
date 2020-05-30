// 互斥锁

// 该代码未学完！！！！
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	// 这里的mutex将同步对state的访问
	// Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁
	var mutex = &sync.Mutex{}
	var ops int64 = 0

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
