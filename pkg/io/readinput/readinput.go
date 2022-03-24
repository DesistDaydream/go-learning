package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("1.使用Scanln函数获取用户的输入")
	ScanlnTest()
	fmt.Println("\n2.使用读取器来获取用户的输入")
	// ReaderTest()
}
func ScanlnTest() {
	var (
		firstName, lastName, s string
		i                      int
		f                      float32
		input                  = "56.12 / 5212 / Go"
		format                 = "%f / %d / %s"
	)

	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}

func ReaderTest() {
	// 使用NewReader函数从系统的标准输入中创建一个读取器。这个读取器会在使用到标准输入的时候，把用户的标准输入的内容存放到缓冲区
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	// 使用ReadString方法作用在读取器变量上，从读取器的缓冲区读取用户的标准输入直到遇上\n，然后连同\n与标准输入的内容一起作为返回值给input变量
	input, err := inputReader.ReadString('\n')
	// 用来判断
	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	// 打印用户的输入的内容
	fmt.Printf("Your name is %s", input)
}
