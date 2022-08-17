package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func filepathWalkArchiving(src string, dstFile *os.File, isGzip bool) (err error) {
	gzipWriter := gzip.NewWriter(dstFile)
	defer gzipWriter.Close()

	// 实例化一个写入器，
	tarWriter := tar.NewWriter(gzipWriter)
	// 检查写入器是否可以成功关闭
	defer func() {
		if err := tarWriter.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// 下面就该开始处理数据了，这里的思路就是递归处理目录及目录下的所有文件和目录
	// 这里可以自己写个递归来处理，不过 Golang 提供了 filepath.Walk 函数，可以很方便的做这个事情
	// 直接将这个函数的处理结果返回就行，需要传给它一个归档源，它就可以自己去处理
	// 我们就只需要去实现我们自己的打包逻辑即可，不需要再去做路径相关的事情
	return filepath.WalkDir(src, func(filePath string, dirEntry fs.DirEntry, err error) error {
		// 因为这个闭包会返回个 error ，所以先要处理一下这个
		if err != nil {
			return err
		}

		// 这里就不需要我们自己再 os.Stat 获取文件信息了了，filepath.Walk 已经做好了，我们直接使用 fileInfo 即可
		fileInfo, _ := dirEntry.Info()
		hdr, err := tar.FileInfoHeader(fileInfo, "")
		if err != nil {
			return err
		}

		// 这里需要处理下 hdr.Name，因为默认文件名是不带路径的，只有一个单纯的名字。
		// 这样打包之后所有文件就会堆在一起，这样就破坏了原本的目录结果，所以需要将文件名，替换为待路径的名称
		// 例如： 将原本 hdr.Name 的 test_dir1 替换成 test_files/test_dir1
		// 这个也很简单，filepath.Walk() 中的回调函数中的 filePath 参数给我们返回来的就是归档源的完整路径
		// 这里的完整路径是相对的，取决于 src 提供的路径，这里面是相对于 scr 如果是目录的话，目录下的每一个文件，都是相对该目录的完整路径
		// hdr.Name = filePath

		// 有一点需要注意，对于 Linux 来说，不能单纯使用 filePath 替换 hdr.Name。如果 src 是绝对路径的话，需要去掉最左侧的 /
		hdr.Name = strings.TrimPrefix(filePath, string(filepath.Separator))

		// 写入文件信息
		if err := tarWriter.WriteHeader(hdr); err != nil {
			return err
		}

		// 判断下文件是否是标准文件，如果不是就不处理了，
		// 如： 目录，这里就只记录了文件信息，不会执行下面的 copy
		if !fileInfo.Mode().IsRegular() {
			return nil
		}

		// 打开文件
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// copy 文件数据到 tarWriter
		n, err := io.Copy(tarWriter, file)
		if err != nil {
			return err
		}

		log.Printf("成功打包 %s ，共写入了 %d 字节的数据\n", filePath, n)

		return nil
	})
}
