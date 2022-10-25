package main

import (
	"bytes"
	"fmt"
	"io"
)

func PrintIOReader(reader io.Reader) {
	fmt.Println(reader)
}

func main() {
	// bytes.Buffer 结构体实现了 io.Reader。
	myBytes := []byte("Hi,DesistDaydream")
	PrintIOReader(bytes.NewBuffer(myBytes))
}
