package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 10 * time.Second

func Control() []time.Time {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	var t []time.Time

	for {
		select {
		// 在此循环期间，如果没有超时，则持续执行
		// 这个地方通常是一个阻塞的操作，由 Channel 持续从其他地方获取数据
		// 直到 ctx.Done 时，将从 Channel 中获取到的数据 return 出去
		case TimeData := <-time.After(3 * time.Second):
			fmt.Println(TimeData)
			t = append(t, TimeData)
		// 如果超时，则直接返回
		case <-ctx.Done():
			return t
		}
	}
}

func main() {
	fmt.Println(Control())
}
