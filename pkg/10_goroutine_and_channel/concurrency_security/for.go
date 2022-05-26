package main

import (
	"log"
	"sync"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	var wg sync.WaitGroup
	defer wg.Wait()

	for _, value := range values {
		wg.Add(1)
		// 错误的写法
		// 这种写法由于闭包只是绑定到这个 value 变量上，并没有被保存到 goroutine 栈中，
		// 所以以上代码极有可能运行的结果都输出为切片的最后一个元素。
		// 因为这样写会导致 for 循环结束后才执行 goroutine 多线程操作，这时候v alue 值只指向了最后一个元素。
		// 这样的结果不是我们所希望的，而且还会产生并发的资源抢占冲突所以是非常不推荐这样写的。
		go func() {
			defer wg.Done()
			log.Printf("错误的写法:%v\n", value)
		}()

		// 正确的写法
		go func(value int) {
			defer wg.Done()
			log.Printf("正确的写法:%v\n", value)
		}(value)
	}
}
