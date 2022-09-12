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

// Marshl 将其他格式的数据转换为 JSON 数据
func Marshl() {
	// 实例化一个结构体数据
	m := Message{"DesistDaydream", "Hello", 1294706395881547000}
	// 将结构体数据编码为 JSON 格式的数据
	b, _ := json.Marshal(m)

	fmt.Println(string(b))
}

// Unmarshl 将 JSON 数据转换为其他格式的数据
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
	var m *Message

	// 将 JSON 格式的数据解码为结构体数据
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(m.Name, m.Body, m.Time)
}

func main() {
	Marshl()
	Unmarshl()
}
