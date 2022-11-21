package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	i := 0
	var wg sync.WaitGroup
	defer wg.Wait()

	// 并发控制。通过一个带缓冲的通道，让并发数量限制在10个
	concurrencyControl := make(chan bool, 10)
	for {
		concurrencyControl <- true
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println(i)
			i += 1
			time.Sleep(1 * time.Second)

			// 每一个 goroutine 执行完毕后，都需要从通道中取出一个数据，相当于释放了通道中一个容量，也可以说是释放了一个并发资源。
			// 如果不取出数据，那么通道中的数据就会一直增加，直到超过缓冲区的大小，从而导致死锁。报错：
			// fatal error: all goroutines are asleep - deadlock!
			<-concurrencyControl
		}()

		if i == 100 {
			return
		}
	}
}
