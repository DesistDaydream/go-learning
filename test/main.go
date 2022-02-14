package main

import (
	"log"
	"time"
)

func main() {
	var data int = 1

	go func() {
		data = 2
		log.Printf("2号协程的值：%v\n", data)
	}()

	go func() {
		data = 3
		log.Printf("3号协程的值：%v\n", data)
	}()

	log.Printf("协程外的值：%v\n", data)

	time.Sleep(2 * time.Second)
}
