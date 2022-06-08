package main

import (
	"embed"
	"fmt"
)

var (
	//go:embed hello.txt
	s string
	//go:embed hello.txt
	b []byte
	//go:embed hello.txt
	f embed.FS
)

func main() {
	fmt.Println(s)
	fmt.Println(string(b))
	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))
}
