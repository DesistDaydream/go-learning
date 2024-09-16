package buffered_io

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

// Attention: Go 按行读取数据的时候有很多需要注意的事项！！！ https://aimuke.github.io/go/2020/06/18/go-readline/

// bufio 包包装了 io 包中 io.Reader 和 io.Writer。
// bufio 包中的各种读/写方法，可以利用 I/O 缓存，减少系统调用，提高性能，处理复杂的 I/O 操作。
// - bufio.Reader
// - bufio.Writer
// - bufio.ReadWriter
// - bufio.Scanner

// Scanner
// Warning: bufio.Scanner 读取的行是有长度限制的，默认是由一个不可更改的变量 MaxScanTokenSize = 64 * 1024 决定的
// 并且该值是固定的，不能被修改的。如果需要读取内容中一行的内容过长，需要使用 bufio.Reader
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
// Warning: bufio.Reader 虽然可以设置缓冲区大小（默认缓冲区 4096 字节）以解决 bufio.Scanner 读取行长度限制的问题。
// 但是 ReadString('\n') 方法读取的内容如果最后一行没有换行符，则不会读取到最后一行
func BufioReader(inputSource io.Reader) {
	fmt.Println("使用 bufio.Reader 读取输入源的内容: ")
	reader := bufio.NewReader(inputSource)

	// 逐行读取文件内容并输出
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print(line) // 是否要在最后一行没有换行符时输出最后一行？
				break
			}
			log.Printf("读取异常: %v", err)
			break
		}
		fmt.Print(line)
	}
}
