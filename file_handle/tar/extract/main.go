package main

func main() {
	// 待提取的文件
	var extractSrc = "./test_files/tets_tar.tar"
	// // 提取后保存的路径，不写就是解压到当前目录
	var extractDst = "./file_handle/tar_dir/"

	Extracting(extractDst, extractSrc)
}
