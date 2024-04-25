package main

import (
	"bufio"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

var srcFile string = "./test_files/test.txt"

// var srcFile string = `F:\Documents\GitHub\Golang\testFile\test.txt`

func main() {
	var action string
	flag.StringVar(&action, "action", "", "")
	flag.Parse()

	switch action {
	case "readfile":
		ReadFileWithByte()
	case "readfile2":
		ReadFileWithBufioReader()
	case "readfile22":
		ReadFileWithBufioScanner()
	case "readfile3":
		ReadFileConvertRowsAndColumns()
	case "readcompress":
		ReadCompress()
	case "writefilebuffer":
		WriteFileBuffer()
	case "copyfile":
		CopyFile()
	}
}

// 最基本的读取读取的方式
func ReadFileDemo(srcFile string) {
	fileByte, _ := os.ReadFile(srcFile)
	fmt.Println("文件内容为：", string(fileByte))
}

// 最基本的文件写入的方式
func WriteFileDemo(dstFile string) {
	os.WriteFile(dstFile, []byte("Hello DesistDaydream"), 0666)
	fileByte, _ := os.ReadFile(dstFile)
	fmt.Printf("写入 %v 文件的内容：%v", dstFile, string(fileByte))
}

// 最基本的文件读取与文件写入的结合示例
func RWFile() {
	inputFile := srcFile
	outputFile := "../testFile/test_copy.txt"
	buf, err := os.ReadFile(inputFile)
	// buf变量类型为[]uint8,也叫[]byte
	fmt.Println(reflect.TypeOf(buf), "\n", buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf)) // 使用string()函数把buf变量的值转变成字符串让人类可读
	// 使用WriteFile()函数把变buf变量中的内容复制到outputFile,若无文件则创建
	os.WriteFile(outputFile, buf, 0644) // oct, not hex
}

// 利用 []byte 处理通过 os.Open 打开的文件内容
func ReadFileWithByte() {
	// 步骤概述：获取文件描述符(简写为FD)，通过FD读取文件放到变量中，然后从变量中输出文件中的内容。
	// 第一步:使用os包中的Open()函数来打开指定文件,返回值为该文件的FD
	fileDescriptor, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer fileDescriptor.Close()
	fmt.Println("打开文件的 FD 为：", fileDescriptor.Fd())
	// 第二步，初始化一个切片变量，使用Read方法，读取Read方法作用的fd变量里的内容，并赋值给括号内的data变量，返回值为指定变量的字节数
	// Read函数详见https://golang.org/pkg/os/#File.Read
	data := make([]byte, 168)
	count, err := fileDescriptor.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// 第三步，使用string函数把byte类型的数据转换成字符型，以便数据可以让人类看懂
	fmt.Printf("比特计数为：%d\n文件内容为：\n%s", count, string(data))
}

// 利用 bufio.Reader 处理通过 os.Open 打开的文件内容
func ReadFileWithBufioReader() {
	// os.File 结构体实现了 io.Reader 接口，可以作为 bufio.NewReader() 函数的参数
	// 注意，很多用于处理 I/O、读写文件 的函数一般都使用 io.Reader 和 io.Writer 接口作为参数，
	// 所以 os.Open 实例化的 File 结构体可以当作很多用来处理 I/O 的函数的参数
	fileDescriptor, err := os.Open(srcFile)
	if err != nil {
		log.Fatalln("文件不存在或出现问题")
	}
	defer fileDescriptor.Close()

	// 使用读取器函数 NewReader() 通过 FD 把文件内容缓存到一个变量中
	// https://golang.org/pkg/bufio/#NewReader
	inputReader := bufio.NewReader(fileDescriptor)
	// 使用 bufio 包中的 Readstring 方法作用在之前初始化的结构体变量上，输出换行符之前的内容，并用无限 for 循环逐行输出，直到最后一行
	// https://golang.org/pkg/bufio/#Reader.ReadString
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Printf("The input was: %s", inputString)
	}
}

// 利用 bufio.NewScanner 处理通过 os.Open 打开的文件内容
func ReadFileWithBufioScanner() {
	fileDescriptor, err := os.Open(srcFile)
	if err != nil {
		log.Fatalln("文件不存在或出现问题")
	}
	defer fileDescriptor.Close()

	// 使用 FD 实例化一个 *bufio.Scanner，通过 Scanner 处理以 ScanLines 定义的分隔符的每部分内容
	// ScanLines 默认是换行符，也就是说，通过 Scanner 的方法都是逐行处理文件内容的。
	// 就像 Scanner 名字似的，扫描仪，扫描仪在扫描文件的时候，也是一行一行扫描的
	scanner := bufio.NewScanner(fileDescriptor)
	// 通过循环逐行处理文件
	for scanner.Scan() {
		// scanner.Text() 返回每行的字符数据
		fmt.Printf("The input was: %s\n", scanner.Text())
	}
}

// ReadFileConvertRowsAndColumns 函数用于行与列互相转换
// 列的数量必须相同，否则切片变量为空
func ReadFileConvertRowsAndColumns() {
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

func ReadCompress() {
	fName := "MyFile.gz"
	var r *bufio.Reader
	fi, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName,
			err)
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}

// WriteFileBuffer 如果您经常将少量数据写入文件，则会降低程序的性能。每次写入都是一个代价高昂的系统调用，如果您不需要立即更新文件，最好将这些小写入归为一个。
// 为此，我们可以使用bufio.Writer结构。它的写入函数不会直接将数据保存到文件中，
// 而是一直保存到下面的缓冲区已满（默认大小为 4096 字节）或Flush()调用该方法。所以一定要Flush()在写入完成后调用，将剩余的数据保存到文件中。
func WriteFileBuffer() {
	var lines = []string{
		"Go",
		"is",
		"the",
		"best",
		"programming",
		"language",
		"in",
		"the",
		"world",
	}
	// create file
	f, err := os.Create(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// 实例化一个 buffer
	buffer := bufio.NewWriter(f)

	for _, line := range lines {
		_, err := buffer.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// 刷新已经缓冲的数据到文件中
	if err := buffer.Flush(); err != nil {
		log.Fatal(err)
	}
}

func CopyFile() {
	dstName := "../testFile/target.txt"
	// 使用Go自带的Args变量获取命令参数。在使用该程序时候，后面加上参数即可把参数赋值给变量srcName。
	// os.Args变量
	srcName := os.Args[1]
	src, err := os.Open(srcName) //打开源文件并获取文件描述符
	if err != nil {
		fmt.Println("无法打开文件")
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644) //打开目标文件并获取文件描述符，如果没有文件则创建，打开的文件`只写`
	if err != nil {
		fmt.Println("无法打开文件")
		return
	}
	defer dst.Close()
	// 通过文件描述符把源文件内容拷贝到目标文件
	io.Copy(dst, src)
}
