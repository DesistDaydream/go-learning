package main

import (
	"fmt"
	"os"
)

// ReadFileConvertRowsAndColumns 函数用于行与列互相转换
// 列的数量必须相同，否则切片变量为空
func ReadFileConvertRowsAndColumns() {
	var srcFile string = "./test_files/test.txt"

	fileDescriptor, err := os.Open(srcFile)
	if err != nil {
		panic(err)
	}
	defer fileDescriptor.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(fileDescriptor, &v1, &v2, &v3)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
