// 通道的遍历
package main

import (
	"fmt"
)

func channelRangeOver() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// 因为之前close了通道，所以下面的迭代会在接收完两个值之后结束
	// 如果没有关闭通道，将会在这个循环中继续阻塞这行，等待接收第三个值
	// 如果没关闭通道，在编译执行的时候会报错
	for elem := range queue {
		fmt.Println(elem)
	}
}
// 注意：一个非空的通道也是可以关闭的，但是在通道中剩下的值仍然可以被接收