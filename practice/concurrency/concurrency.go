package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	// var wg sync.WaitGroup
	// defer wg.Wait()

	// 并发控制，通过一个带缓冲的通道，让并发数量限制在10个
	cc := make(chan bool, 10)
	for {
		cc <- true
		// wg.Add(1)
		go func() {
			fmt.Println(i)
			i += 1
			time.Sleep(1 * time.Second)
			<-cc
		}()
		// wg.Done()
		if i == 1000 {
			return
		}
	}
}
