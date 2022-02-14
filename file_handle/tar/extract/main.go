package main

func main() {
	// 待提取的文件
	var extractSrc = "file_handle/tar_dir/tets_tar.tar.gz"
	// // 提取后保存的路径，不写就是解压到当前目录
	var extractDst = "file_handle/tar_dir/"

	Extracting(extractSrc, extractDst)
}
