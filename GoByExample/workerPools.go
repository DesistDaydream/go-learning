// 工作池，使用Go协程和通道实现工作池
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 为了使用worker工作池并收集其结果，需要2个通道
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 启动3个worker，初始是阻塞的，因为还没有传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 发送了9个jobs，然后关闭这些通道来表示这些就是所有的任务了
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	// 收集所有这些任务的返回值
	for a := 1; a <= 9; a++ {
		<-results
	}
}
// 这个程序显示9个任务被多个worker执行。
// 整个程序处理所有的任务仅执行3s而不是9s，是因为3个worker是并行的。
