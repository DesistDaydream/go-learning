package main

import (
	"embed"
	"fmt"
	"net/http"
)

var (
	//go:embed hello.txt
	s string
	//go:embed hello.txt
	b []byte
	//go:embed hello.txt
	f embed.FS
)

func httpEmbed() {
	http.Handle("/", http.FileServer(http.FS(f)))
	http.ListenAndServe(":8080", nil)
}

func main() {
	// 基础使用
	// 将嵌入的文件转为字符串
	fmt.Println(s)
	// 将嵌入的文件转为字节数组
	fmt.Println(string(b))
	// 将嵌入的文件转为文件系统
	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))

	// 常用于 Web 服务，以嵌入 HTML 等静态文件
	httpEmbed()
}
