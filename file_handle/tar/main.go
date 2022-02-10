package main

import "log"

func main() {
	// 待打包的目录
	var archiveSrc = "practice"
	// 打包后的文件名
	var archiveDst = "./file_handle/tar_dir/test_tar.tar.gz"

	if err := Archiving(archiveSrc, archiveDst); err != nil {
		log.Fatalln(err)
	}

	// 待提取的文件
	var extractSrc = "./file_handle/tar_dir/test_tar.tar.gz"
	// 提取后保存的路径，不写就是解压到当前目录
	var extractDst = "./file_handle/tar_dir/"

	Extracting(extractDst, extractSrc)
}
