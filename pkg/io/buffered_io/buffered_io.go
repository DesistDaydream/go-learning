package buffered_io

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

// bufio 包包装了 io 包中 io.Reader 和 io.Writer。
// bufio 包中的各种读/写方法，可以利用 I/O 缓存，减少系统调用，提高性能，处理复杂的 I/O 操作。
// - bufio.Reader
// - bufio.Writer
// - bufio.ReadWriter
// - bufio.Scanner

// Scaaner
func BufioScanner(inputSource io.Reader) {
	fmt.Println("使用 bufio.Scanner 读取输入源的内容:")
	scanner := bufio.NewScanner(inputSource)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("扫描文件时发生错误: %v", err)
	}
}

// Reader
func BufioReader(inputSource io.Reader) {
	fmt.Println("使用 bufio.Reader 读取输入源的内容: ")
	reader := bufio.NewReader(inputSource)

	// 逐行读取文件内容并输出
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print(line)
	}
}
