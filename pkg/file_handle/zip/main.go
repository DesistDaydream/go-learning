package main

import "log"

func Archiving() {
	// 源档案（准备压缩的文件或目录）
	var src = "log"
	// 目标文件，压缩后的文件
	var dst = "log.zip"

	if err := Zip(dst, src); err != nil {
		log.Fatalln(err)
	}
}

func Extracting() {
	// 压缩包
	var src = "log.zip"
	// 解压后保存的位置，为空表示当前目录
	var dst = ""

	if err := UnZip(dst, src); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	Archiving()
	Extracting()
}
