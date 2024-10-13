package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用带缓冲的通道来控制并发数量
func demo() {
	var i int = 0

	var wg sync.WaitGroup
	defer wg.Wait()

	// 一、控制并发数量。通过一个带缓冲的通道，让并发数量限制在10个
	// TODO: 使用 bool 或 struct{} 作为通道的类型有什么区别？`concurrencyControl := make(chan struct{}, 10)`
	concurrencyControl := make(chan bool, 10)
	for i < 100 {
		// 二、获取并发资源
		// 若用 struct{} 作为通道的类型，则需要使用 `concurrencyControl <- struct{}{}`
		concurrencyControl <- true

		wg.Add(1)

		go func(num int) {
			defer wg.Done()
			// 三、释放并发资源
			// 每一个 goroutine 执行完毕后，都需要从通道中取出一个数据，相当于释放了通道中一个容量，也可以说是释放了一个并发资源。
			// 如果不取出数据，那么通道中的数据就会一直增加，直到超过缓冲区的大小，从而导致死锁。
			// 死锁报错: fatal error: all goroutines are asleep - deadlock!
			defer func() { <-concurrencyControl }()
			// 也可以直接在函数结尾使用 `<-concurrencyControl` 这种表达式，效果是一样的。不过代码可读性差，而且函数结尾的位置不好确定

			fmt.Println(num)

			// 加这个只是为了让程序执行时间更长一些，方便观察。否则开了并发瞬间就都完成了。
			time.Sleep(1 * time.Second)
		}(i)

		i++
	}
}

func main() {
	demo()
	// 这里仅仅演示是对并发控制进行一个演示。主要是演示控制并发数量的，并没有解释 WaitGroup 相关代码的含义
	// 除了使用 WaitGroup 控制并发外，还可以使用 context。
	// 更多并发控制描述详见 https://github.com/DesistDaydream/notes-learning/blob/main/content/zh-cn/docs/2.%E7%BC%96%E7%A8%8B/%E9%AB%98%E7%BA%A7%E7%BC%96%E7%A8%8B%E8%AF%AD%E8%A8%80/Go/Go%20%E8%A7%84%E8%8C%83%E4%B8%8E%E6%A0%87%E5%87%86%E5%BA%93/Goroutine%20AND%20Channel/%E5%B9%B6%E5%8F%91%E6%8E%A7%E5%88%B6.md
}
