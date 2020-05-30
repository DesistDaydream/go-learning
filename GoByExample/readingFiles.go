// 读文件

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)



// 读取文件需要进行错误检查，这个方法可以精简下面的错误检查过程
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 基本的读取文件，并把文件内容作为返回值赋值给变量data，该返回值类型为`[]byte`
	// 最后使用string()把byte类型的数据转换成字符串以便人类阅读
	data, err := ioutil.ReadFile("../testFile/test.txt")
	check(err)
	fmt.Println(string(data))

	// 如果想要读取一个文件的部分内容，或者逐行读取或者其余对这个文件进行更多的控制
	// 则需要先打开这个文件，`op.Open`函数的返回值为对应文件的描述符，后续对文件的处理，都通过文件描述符来进行
	f, err := os.Open("../testFile/test.txt")
	check(err)
	// 在程序结束前，关闭这个文件以节省资源
	defer f.Close()

	fmt.Printf("\n测试一：从文件开始位置读取一些字节,这里最多读取5个字节\n")
	test1(f)

	fmt.Printf("\n测试二：使用`Seek`从文件已知的位置开始进行读取\n")
	test2(f)

	fmt.Printf("\n测试三：io包提供了一些可以帮助我们进行文件读取的函数，测试二的读取可以使用ReadAtLeast函数得到一个更健壮的实现\n")
	test3(f)

	fmt.Printf("\n测试四：没有内置的回转支持(built-in rewind)，但是使用Seek(0,0)实现\n")
	_, err = f.Seek(0, 0)
	check(err)

	fmt.Printf("\n测试五：bufio包实现了带缓冲的读取，这不仅对于很多小的读取啊哦做能够提升性能，也提供了很多附加的读取函数\n")
	test5(f)
	// 注意，测试五与测试六不可同时使用，否则报错
	fmt.Printf("\n测试六：使用bufio.Scan来逐行读取文件\n")
	test6(f)
}





// 从文件开始位置读取一些字节,这里最多读取5个字节
func test1(f *os.File) {
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}

// Seek函数设置下一次读/写的位置，第一个参数为相对偏移量，第二个参数决定相对位置；0为相对文件开头，1为相对当前位置，2为相对文件结尾,返回新的偏移量(相对开头)和可能得错误
// 该例子是从第6个字符后面开始读取，读取2个字节
func test2(f *os.File) {
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
}

// 测试二的读取可以使用ReadAtLeast函数得到一个更健壮的实现
func test3(f *os.File) {
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
}

// `NewReader`函数会创建一个具有默认大小缓冲且从参数中指定的文件读取的读取器，返回值为`*Reader`
// `Peek`函数返回输入流的参数指定的后n个字节，该例子为后5个
func test5(f *os.File) {
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}

func test6(f *os.File) {
	fmt.Println(f)
	var count int = 0
	r5 := bufio.NewScanner(f)
	fmt.Println(r5)
	for r5.Scan() { 
		count++
		fmt.Printf("the line %d: %s\n", count, r5.Text())
	}
}


