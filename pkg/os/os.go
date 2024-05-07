package main

import (
	"fmt"
	"os"
	"path/filepath"

	"log"
)

func traversalDir() {
	dir := "./pkg/os"
	fmt.Println("for range 遍历目录")
	// for range 遍历目录
	// 仅读取目录下一层的文件，不会递归遍历子目录
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		log.Printf("%s", file.Name())
	}

	fmt.Println("filepath.Walk 遍历目录")
	// filepath.Walk 遍历目录
	// 递归遍历目录及其子目录
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		log.Printf("%s", path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// for range 与 filepath.Walk 遍历目录的示例
	// 主要区别在于是否递归遍历子目录
	traversalDir()
}
