package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

func HelloWorld(srcFile io.Reader, dstFile string, isGzip bool) {
	// 1. 实例化一个 tar 读取器
	tr := tar.NewReader(srcFile)
	if isGzip {
		gr, _ := gzip.NewReader(srcFile)
		defer gr.Close()
		tr = tar.NewReader(gr)
	}

	// 2. 获取文件头信息
	hdr, _ := tr.Next()

	// 根据文件头信息判断文件类型
	switch hdr.Typeflag {
	case tar.TypeDir:
		os.MkdirAll(dstFile+hdr.Name, 0775)
	case tar.TypeReg:
		// 如果是文件，就写入
	}

	// 注意：与归档类似，这里的将要提取的  srcFile 虽然是一个目录，但是在提取的过程中，只会将目录本身提取出来
	// 也就是说，目录下的所有内容都不会自动处理，所以我们要自己递归目录，逐一将文件提取出来
}

func main() {
	// 待提取的文件
	var extractSrc = "file_handle/tar_dir/test_tar.tar.gz"
	// 提取后保存的路径，不写就是解压到当前目录
	var extractDst = "file_handle/tar_dir/"
	var isGzip = true

	// 打开准备提取的 tar 包
	file, _ := os.Open(extractSrc)
	defer file.Close()

	HelloWorld(file, extractDst, isGzip)

	// Extracting(file, extractDst, isGzip)
}
