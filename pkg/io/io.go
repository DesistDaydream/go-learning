package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/DesistDaydream/go-learning/pkg/io/buffered_io"
)

var inputSource string = "./test_files/test.txt"

// 注意: 要想理解 io，一定要先理解 Go 中的 Interface 概念
// 注意: 要想理解 io，一定要先理解 Go 中的 Interface 概念
// 注意: 要想理解 io，一定要先理解 Go 中的 Interface 概念

// io.Reader、io.Writer 是 io 包中的接口，用于处理 I/O 操作。
// 所有实现了 io.Reader 接口的类型都可以作为输入源，例如 打开的文件、建立的网络连接、etc.
// 所有实现了 io.Writer 接口的类型都可以作为输出目标，例如 打开的文件、建立的网络连接、etc.
// 很多方法方法和函数都会使用 io.Reader 或 io.Writer 作为参数的类型，以便执行读写逻辑，I/O 逻辑是非常基本操作，很多地方都会用到

// I/O 示例
// 实例中展示了多种实现了 io.Reader 接口的类型，可以当作 I/O 输入源使用。
func IOReadDemo() {
	// strings.Reader 实现了 io.Reader
	inputSource := strings.NewReader("strings.Reader 类型实例化时使用了这段字符串，该类型实现了 io.Reader")
	useReader(inputSource)

	// bytes.Buffer 实现了 io.Reader
	buffer := bytes.NewBufferString("bytes.Buffer 类型实例化时使用了这段字符串，该类型实现了 io.Reader")
	useReader(buffer)

	// os.File 实现了 io.Reader
	file, _ := os.Open("test_files/test.txt")
	useReader(file)
	file.Close()

	// HTTP 响应体本身就是 io.ReadCloser，实现了 io.Reader，顺带实现了 io.Closer
	resp, _ := http.Get("https://example.com/")
	useReader(resp.Body)

	// net.Conn 实现了 io.Reader
	// 注意：这个示例会报错，必须有可以连接的 TCP 才可以
	conn, _ := net.DialTimeout("tcp", "0.0.0.0:23", 2*time.Second)
	useReader(conn)
	conn.Close()
}

// 这里是一个使用 io.Reader 的示例，**只要传递进来的输入源的类型实现了 io.Reader**，即可从输入源中读取出数据
func useReader(inputSource io.Reader) {
	// 适当大小的缓冲区。
	bufferSize := 1024
	buf := make([]byte, bufferSize)
	// var buf [20]byte // 也可以使用这种方式声明

	fmt.Printf("检查 inputSource 类型: %v\n", reflect.TypeOf(inputSource))

	// for 循环中，每次从 inputSource 中读取 bufferSize 大小的数据，并将其存储到 buf 中
	// 假如 inputSource 有 120 bytes，但是缓冲器大小只有 100 bytes，那么输出结果就是 2 行，第一行是 100 bytes，第二行是 20 bytes
	// 根据 byte 分隔一串字符会有一个问题，比如某个字符占用 2 两个字节的话，此时这个字符就被会分割成两部分，解码时就会变成乱码。
	// 一般 bufferSize 都是 2^n 的形式，这样可以保证每次读取的字节数都是固定的，不会出现乱码
	for {
		n, err := inputSource.Read(buf)
		// fmt.Println("检查 Read 返回的错误: ", err)
		// 文件或者纯文本的结尾一般都有 EOF，用以表示这段内容已经结束，当 err 中为 EOF 则表明这部分内容已经读取完毕，可以退出循环
		// 注意:
		// 网络连接没有 EOF，因为网络连接通常会持续得等待连接中的数据到达（毕竟不知道数据是否发送完毕，除非是 HTTP 那种有明确结构得，或者对方直接关闭连接）。
		// 所以若是 inputSource 是 net.Conn 类型，默认不会有 ERROR，毕竟数据还没有发送完并一直等待，也就不可能会有 EOF 相关的返回，该循环也不会退出
		// net.Conn 如果想要想要知道数据是否已经读取完毕，需要自行使用 net.Conn.SetReadDeadline 设置读取超时的时间点，
		// 若到达时间点后还是没有数据，则会返回超时的错误。此时读取到的 n 为 0。
		// 从 Go 标准库 net/net.go 中可以看到 Error 接口只有一个 Timeout() 方法
		// 其中有个 Temporary() 方法已经注释弃用，注释中提到了绝大多数所谓的临时错误都是超时导致的。所以想要判断一个 net.Conn 是否已经读取完毕，主要就是靠自己设定的超时错误。
		if err != nil {
			if err == io.EOF {
				log.Println("全部读取完成")
				break
			}
			log.Fatalf("读取错误：%v", err)
		}
		fmt.Printf("从 inputSource 读取到的字符串: %q\n", string(buf[:n]))
	}
	fmt.Printf("\n")
	// 如果想要逐行读取，这需要在 for 循环中逐个字节地读取 inputSource，并在遇到换行符时停止读取并处理当前行。
	// 可以使用下面的 useReaderLineByLine() 函数
	// 或者使用 bufio 标准库中的相关方法，例如 bufio.Scanner、bufio.Reader 等，这些方法可以更方便地处理 I/O 操作。

	// 注意: io.ReadAll 是对 Reader.Read() 的封装，用于一次性读取所有数据，并返回读取到的数据
	// 但是，无法处理某些特定类型的 io.Reader，比如 net.conn，因为 net.conn 这种网络连接形式的 io.Reader 可能会阻塞导致 ReadAll 一直等待更多的数据到达
	// 所以需要采用逐块读取的方式，读取一部分数据，然后处理它，接着继续读取下一部分，直到读取完成或发生错误。
	// 除了使用上面的 Reader.Read() 方法外，还可以使用 bufio 包中的相关方法，例如 bufio.Scanner、bufio.Reader 等，这些方法可以更方便地处理 I/O 操作。
	// inputBytes, err := io.ReadAll(inputSource)
	// if err != nil {
	// 	log.Fatalf("%v", err)
	// }
	// fmt.Println(string(inputBytes))
}

// 使用 io.Reader 逐行读取
func useReaderLineByLine(inputSource io.Reader) {
	var (
		buf  [1]byte // 逐字节读取是临时存储每隔字节信息的 buffer
		line []byte  // 保存当前行内容的切片。每次读取一个字节，将其追加到 line 切片末尾，直到遇到换行符 '\n'，然后将 line 切片打印出来
	)

	for {
		n, err := inputSource.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				log.Println("全部读取完成")
				break
			}
			log.Fatalf("读取错误：%v", err)
		}
		if n == 0 {
			continue
		}

		// 将当前字节追加到行内容切片
		line = append(line, buf[0])

		// 如果当前字节是换行符，则处理当前行
		if buf[0] == '\n' {
			fmt.Printf("%s", line) // 处理当前行
			line = nil             // 清空当前行内容切片，准备处理下一行
		}
	}

	// 处理剩余的内容
	if len(line) > 0 {
		fmt.Printf("%s", line)
	}
}

func main() {
	demo := flag.String("demo", "bufio", "使用哪个示例")
	flag.Parse()

	switch *demo {
	case "io":
		// 简单 I/O Read 示例
		IOReadDemo()
	case "bufio":
		// bufio 库示例
		file, err := os.Open(inputSource)
		if err != nil {
			log.Fatalf("打开文件失败: %v", err)
			return
		}
		defer file.Close()

		// 通过 bufio.Scanner 读取数据
		buffered_io.BufioScanner(file)
		// 将读取位置重置到文件开头，避免 BufioReader 读取不到数据。
		// 因为 buffio 库相关的方法会将 io.Reader 类型的变量读取位置标记为文件末尾，下次再读取时，会直接从文件末尾开始读取，导致读取不到数据。
		file.Seek(0, 0)
		// 通过 bufio.Reader 读取数据
		buffered_io.BufioReader(file)
	}
}
