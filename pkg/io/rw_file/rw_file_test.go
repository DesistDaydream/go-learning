package main

import (
	"testing"
)

var SrcFile string = "../../../test_files/test.txt"
var DstFile string = "../../../test_files/sample_write_file.txt"

func TestReadFileDemo(t *testing.T) {
	type args struct {
		srcFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "最简单的读取文件的方法",
			args: args{
				srcFile: SrcFile,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadFileDemo(SrcFile)
		})
	}
}

func TestWriteFileDemo(t *testing.T) {
	type args struct {
		dstFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "最简单的写入文件的方法",
			args: args{
				dstFile: DstFile,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteFileDemo(tt.args.dstFile)
		})
	}
}
