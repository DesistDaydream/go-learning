package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Message 一条消息应该具有的属性
type Message struct {
	Name string
	Body string
	Time int64
}

// Unmarshl 将 JSON 数据解码为 Go 程序可以理解的数据类型（比如 struct、slice 等等）
// 字符串数据转为结构体数据
func Unmarshl() {
	// 这是 JSON 格式的数据
	jsonStr := `
	{
		"Name":"DesistDaydream",
		"Body":"Hello",
		"Time":1294706395881547000
	}`
	// 声明一个变量，用来存放 JSON 解码后的数据
	// 注意：若要将 JSON 的数据保存到 struct 中，则需要保证在 struct 中的每个字段，都对应 JSON 数据中的每个 key
	var m Message

	// 将 JSON 格式的数据解码为结构体数据
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(m.Name, m.Body, m.Time)
}

// Marshl 将 Go 程序可以理解的数据类型（比如 struct、slice 等等）编码为 JSON 数据
// 结构体数据转为字符串数据
func Marshl() {
	// 实例化一个结构体数据
	m := &Message{"DesistDaydream", "Hello", 1294706395881547000}

	// 将结构体数据编码为 JSON 格式的数据
	b, _ := json.Marshal(m)

	fmt.Println(string(b))
}

func main() {
	// 编码和解码的对象是相对于原始字符串的 JSON 数据。所以
	// 编码就是将其他数据编码成 JSON 数据
	// 解码就是将 JSON 数据解码为其他数据
	Marshl()
	Unmarshl()
}
