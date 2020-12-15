// 退出状态
// 使用os.Exit来立即退出该Go程序，并携带执行的退出状态码
package main

import (
	"fmt"
	"os"
)

func exitHandling() {
	defer fmt.Println("!")
	// 当使用os.Exit时，defer将不会执行
	os.Exit(3)
}
