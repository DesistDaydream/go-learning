package main

import "fmt"

func main() {
    fmt.Println("Starting the program")
    // panic可以直接从代码初始化
    // panic()函数产生一个终止程序的运行时错误，并打印给该函数传递的参数。
    panic("A severe error occurred: stopping the program!")
    fmt.Println("Ending the program")
}