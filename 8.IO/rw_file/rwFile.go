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

var srcFile string = "./test_file/test.txt"

// var srcFile string = `F:\Documents\GitHub\Golang\testFile\test.txt`
var useage string = `可用的值有：
samplereadfile:最简单的读取文件的方法
samplewritefile:最简单的写入文件的方法
readfile:读取一个文件的内容并输出
readfile2:逐行读取一个文件的内容并逐行输出的方法1
readfile22:逐行读取一个文件的内容并逐行输出的方法2
readfile3:行与列互换
rwfile:读取一个文件的内容并复制到一个新文件中
readcompress:读取压缩文件
writefilebuffer:使用缓冲写入文件。
copyfile:通过命令行参数把一个文件复制到另一个文件
`

func main() {
	var action string
	flag.StringVar(&action, "action", "", useage)
	flag.Parse()

	switch action {
	case "samplereadfile":
		SampleReadFile()
	case "samplewritefile":
		SampleWriteFile()
	case "readfile":
		ReadFile()
	case "readfile2":
		ReadFile2()
	case "readfile22":
		ReadFile22()
	case "readfile3":
		ReadFile3()
	case "rwfile":
		RWFile()
	case "readcompress":
		ReadCompress()
	case "writefilebuffer":
		WriteFileBuffer()
	case "copyfile":
		CopyFile()
	}
}

func SampleReadFile() {
	fileByte, _ := os.ReadFile(srcFile)
	fmt.Println("文件内容为：", string(fileByte))
}

func SampleWriteFile() {
	os.WriteFile("./test_file/SampleWriteFile", []byte("Hello DesistDaydream"), 0666)
	fileByte, _ := os.ReadFile("./test_file/SampleWriteFile")
	fmt.Println("写入的文件为：", string(fileByte))
}

func ReadFile() {
	// 以下变量声明仅用来查看其类。因为下面的代码的变量在初始化的时候直接声明加赋值了，没法显式得看出变量的类型
	// var fd *os.File
	// var data []uint8	//uint8与byte类型一样

	// 步骤概述：获取文件描述符(简写为FD)，通过FD读取文件放到变量中，然后从变量中输出文件中的内容。
	// 第一步:使用os包中的Open()函数来打开指定文件,返回值为该文件的FD
	fd, err := os.Open(srcFile)
	fmt.Println("打开文件的 FD 为：", fd.Fd())
	// 下面几行不影响代码功能，可省略的代码，判断文件是否存在，如果不存在会提示
	if err != nil {
		log.Fatal(err)
	}
	// 第二步，初始化一个切片变量，使用Read方法，读取Read方法作用的fd变量里的内容，并赋值给括号内的data变量，返回值为指定变量的字节数
	// Read函数详见https://golang.org/pkg/os/#File.Read
	data := make([]byte, 168)
	count, err := fd.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	// 第三步，使用string函数把byte类型的数据转换成字符型，以便数据可以让人类看懂
	fmt.Printf("比特计数为：%d\n文件内容为：\n%s", count, string(data))
}

func ReadFile2() {
	// var inputFile *os.File
	// var inputReader *bufio.Reader
	// var inputString
	inputFile, inputError := os.Open(srcFile)
	if inputError != nil {
		fmt.Printf("文件不存在或出现问题")
		return // 出错时退出这个函数
	}
	defer inputFile.Close() // 确保正常可以在函数结束前关闭打开的文件

	// 使用读取器函数NewReader()通过FD把文件内容缓存到一个变量中
	// https://golang.org/pkg/bufio/#NewReader
	inputReader := bufio.NewReader(inputFile)
	// 第三步:使用bufio包中的Readstring方法作用在之前初始化的结构体变量上，输出换行符之前的内容，并用无限for循环逐行输出，直到最后一行
	// https://golang.org/pkg/bufio/#Reader.ReadString
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			return
		}
		fmt.Printf("The input was: %s", inputString)
	}
}

func ReadFile22() {
	// var inputFile *os.File
	// var inputReader *bufio.Reader
	// var inputString
	var slices []string
	inputFile, inputError := os.Open(srcFile)
	if inputError != nil {
		fmt.Printf("文件不存在或出现问题")
		return // 出错时退出这个函数
	}
	defer inputFile.Close() // 确保正常可以在函数结束前关闭打开的文件

	// 使用读取器函数NewScanner()通过FD把文件内容缓存到一个变量中
	scanner := bufio.NewScanner(inputFile)
	// 第三步:使用bufio包中的Readstring方法作用在之前初始化的结构体变量上，输出换行符之前的内容，并用无限for循环逐行输出，直到最后一行
	// https://golang.org/pkg/bufio/#Reader.ReadString
	for scanner.Scan() {
		fmt.Printf("The input was: %s\n", scanner.Text())
		// 还可以通过append函数将每一行内容赋值给一个切片，以便后续使用
		slices = append(slices, scanner.Text())
		fmt.Println(slices)
	}
}

// ReadFile3 函数用于行与列互相转换
// 列的数量必须相同，否则切片变量为空
func ReadFile3() {
	file, err := os.Open(srcFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
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
