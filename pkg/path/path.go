package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func Path() {
	// 从一个绝对路径中提取出最后一个元素
	// 在提取最后一个元素之前删除尾部斜杠。如果路径为空，Base 返回 “.”。如果路径完全由斜杠组成，则 Base 返回 “/”。
	fmt.Println(path.Base("/a/b/"))
	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
}

func FilePath() {
	file := "test.txt"
	dir := "test_files"

	// 拼接路径
	filePath := filepath.Join(dir, file)
	fmt.Println(filePath)
}

func main() {
	Path()
	FilePath()
}
