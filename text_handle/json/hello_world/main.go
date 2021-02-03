package main

import (
	"encoding/json"
	"fmt"
)

// Message 一条消息应该具有的属性
type Message struct {
	Name string
	Body string
	Time int64
}

// Marshl 将其他格式的数据转换为 JSON 数据
func Marshl() {
	m := Message{"DesistDaydream", "Hello", 1294706395881547000}
	// 使用 Marshal() 方法，将 m 编码为 b
	b, _ := json.Marshal(m)

	fmt.Println(string(b))
}

func main() {
	Marshl()
}
