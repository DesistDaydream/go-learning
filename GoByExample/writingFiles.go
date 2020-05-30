package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 展示如何写入一个字符串(或者只是一些字节)到一个文件
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("../testFile/testWrite.txt", d1, 0644)
	check(err)

	// 采用0666模式创建一个文件，返回该文件的描述符，如果要创建的文件已存在则会覆盖
	f, err := os.Create("../testFile/testWrite2.txt")
	check(err)
	// 打开文件后习惯性得使用defer调用文件的close操作
	defer f.Close()

	// 可以写入任何想写入的字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes \n", n2)

	// 也可以直接写入字符串
	n3, err := f.WriteString("wrties\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用Sync来将缓冲区的信息写入磁盘中
	f.Sync()

	// 带缓冲的写入器
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes \n", n4)

	// 使用Flush确保所有缓存的操作记录已写入底层写入器
	w.Flush()

}
