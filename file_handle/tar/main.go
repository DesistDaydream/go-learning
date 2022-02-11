package main

import (
	"archive/tar"
	"os"
)

func ArchivingHelloWorld(src, dst string) {
	// 归档大体分为三步

	// 创建归档目标文件，等待归档
	file, _ := os.Create(dst)
	defer file.Close()

	// 1. 实例化一个 tar 写入器，写入器主要用来将数据写入归档文件中，即 file 变量
	// 这里将 file 连接到了 tw.w 属性
	// 当我们将归档源写入 tw.w 属性，就相当于是写入到归档文件中了。
	tw := tar.NewWriter(file)
	// 如果不关闭会造成归档文件不完整，不完整的归档文件打开后会报错
	defer tw.Close()

	// 2. 将归档源的 Header 信息写入归档文件中
	// 获取归档源信息
	fileInfo, _ := os.Stat(src)
	// 通过归档源的文件信息，生成用于填充 tar 的 Header
	hdr, _ := tar.FileInfoHeader(fileInfo, "")
	// 将 Header 写入归档文件中
	tw.WriteHeader(hdr)

	// 3. 将归档源写入到归档文件中
	// 打开归档源，并将 归档源 写入归档文件中
	srcFile, _ := os.ReadFile(src)
	tw.Write(srcFile)

	// 注意：这里的 src 虽然是一个目录，但是在归档的过程中，只会将目录本身写入归档文件
	// 目录下的所有内容都不会自动处理，所以我们要自己递归目录，逐一将文件写入归档文件
	// 也可以使用 Go 标准库中的 filepath.Walk() 函数将目录自动递归。
}

func main() {
	// 归档源，即待归档的目录
	var archiveSrcPath = "test_file"
	// 归档目，即归档后生成的文件
	var archiveDstPath = "test_file/test_tar.tar"

	ArchivingHelloWorld(archiveSrcPath, archiveDstPath)

	// if err := Archiving(archiveSrc, archiveDst); err != nil {
	// 	log.Fatalln(err)
	// }

	// 待提取的文件
	// var extractSrc = "./file_handle/tar_dir/test_tar.tar.gz"
	// // 提取后保存的路径，不写就是解压到当前目录
	// var extractDst = "./file_handle/tar_dir/"

	// Extracting(extractDst, extractSrc)
}
