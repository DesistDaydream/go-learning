package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

var tw *tar.Writer

func HelloWorld(srcFile string, dstFile io.Writer, isGzip bool) {
	// 归档大体分为三步

	// 1. 实例化一个 tar 写入器，写入器主要用来将数据写入归档文件中，即 file 变量
	if !isGzip {
		// 这里将 dstFile 连接到了 tw.w 属性
		// 当我们将归档源写入 tw.w 属性，就相当于是写入到归档文件中了。在代码最后，就会调用 tw.Write(SrcFIieByte) 方法，将归档源写入到 tw.w 中，即写入到的 dstFile 中
		tw = tar.NewWriter(dstFile)
	} else {
		// 将 tar 包使用 gzip 压缩，只需要在 dstFile 和 writer 之前加上一层压缩就行了，和 Linux 的管道的感觉类似
		gzipWriter := gzip.NewWriter(dstFile)
		defer gzipWriter.Close()

		tw = tar.NewWriter(gzipWriter)
	}
	// 如果不关闭会造成归档文件不完整，不完整的归档文件打开后会报错
	defer tw.Close()

	// 2. 将归档源的 Header 信息写入归档文件中
	// 获取归档源信息
	fileInfo, _ := os.Stat(srcFile)
	// 通过归档源的文件信息，生成用于填充 tar 的 Header
	hdr, _ := tar.FileInfoHeader(fileInfo, "")
	// 将 Header 写入归档文件中
	tw.WriteHeader(hdr)

	// 3. 将归档源写入到归档文件中
	// 打开归档源，并将 归档源 写入归档文件中
	srcFileByte, _ := os.ReadFile(srcFile)
	tw.Write(srcFileByte)

	// 注意：这里的 srcFile 虽然是一个目录，但是在归档的过程中，只会将目录本身写入归档文件
	// 也就是说，目录下的所有内容都不会自动处理，所以我们要自己递归目录，逐一将文件写入归档文件
	// 也可以使用 Go 标准库中的 filepath.Walk() 函数将目录自动递归。
}

func main() {
	// 归档源，即待归档的目录
	var archiveSrcPath = "test_files"
	// 归档目，即归档后生成的文件
	// var archiveDstPath = "test_files/test_tar.tar"
	var archiveDstPath = "file_handle/tar_dir/test_tar.tar.gz"
	// 是否使用 Gzip 压缩 tar 归档文件
	isGzip := true

	// 创建归档目标文件，等待归档
	file, _ := os.Create(archiveDstPath)
	defer file.Close()

	// 归档功能的基本介绍
	// HelloWorld(archiveSrcPath, file, isGzip)

	// 通过 filepath.Walk() 函数递归归档
	filepathWalkArchiving(archiveSrcPath, file, isGzip)
}
