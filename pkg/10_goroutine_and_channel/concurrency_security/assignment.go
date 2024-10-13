package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

const gorotineNum int = 100000

func demo() {
	var counter int32 = 0
	var wg sync.WaitGroup

	// 启动 100 个 goroutine
	for i := 0; i < gorotineNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 如果是并发赋值，则会出现数据竞争。最后的结果必然是一个小于 100000 的随机数
			// counter++
			// 通过使用 atomic 包中的 AddInt32 函数，可以保证并发赋值的安全性
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("计数结果: ", counter)
}

func demoWithMutex() {
	var count int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < gorotineNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("利用 Mutex 的计数结果: ", count)
}

func demoWithChannel() {
	var count int
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < gorotineNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1
		}()
	}

	go func() {
		for c := range ch {
			count += c
		}
	}()

	wg.Wait()
	fmt.Println("利用 Channel 的计数结果: ", count)
}

func statFileLines() {
	var wg sync.WaitGroup
	var fileLineCount int32 = 0

	countLine := func(filePath string) {
		defer wg.Done()
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		var lineCount int
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineCount++
		}

		// 这里由于文件少，读取文件时间较长，每次赋值并不一定会产生并发赋值安全问题
		atomic.AddInt32(&fileLineCount, int32(lineCount))
		// fileLineCount = fileLineCount + int32(lineCount)
	}

	err := filepath.Walk("/mnt/d/Projects/DesistDaydream/go-learning/test_files", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(info.Name()) == ".go" || filepath.Ext(info.Name()) == ".txt" {
			wg.Add(1)
			go countLine(path)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("遍历目录失败: %v", err)
	}

	wg.Wait()

	fmt.Println("总行数: ", fileLineCount)
}

func main() {
	// 在 concurrency_control 部分并没有涉及到 **并发赋值** 的行为，都是将赋值行为放在并发逻辑外面进行的。
	// 那如果想要在并发里面进行赋值呢？如果同时对一个变量进行并发赋值，会发生什么？
	statTime(demo) // 21.26ms

	// 若是并发赋值不是简单的计数器，而是需要保证数据一致性的话，可以使用 sync 包提供的 Mutex 或者 RWMutex 来保证数据一致性
	statTime(demoWithMutex) // 23.40ms

	// 还有一种利用 Channel 来保证并发赋值安全的方式
	// 但是，利用 Channel 的方式效率相对较低，需要花费更长的时间
	// 相对的，好处是更加灵活，可以在遍历 Channel 时处理更复杂的场景
	statTime(demoWithChannel) // 131.53ms

	// 并发时赋值的一个典型场景：同时统计多个文件的行数
	statFileLines()
}

func statTime(fn func()) {
	start := time.Now()
	fn()
	fmt.Printf("函数执行时间: %s\n", time.Since(start))
	fmt.Println("=======================")
}
