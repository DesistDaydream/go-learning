package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// 两种解码的区别
func UnmarshalAndDecode() {
	jsonStr := `
	{
		"Name":"DesistDaydream",
		"Body":"Hello",
		"Time":1294706395881547000
	}`
	type message struct {
		Name string
		Body string
		Time int64
	}
	var m message

	// 使用 Unmarshl()
	json.Unmarshal([]byte(jsonStr), &m)
	fmt.Println(m)

	// 使用 Decode()
	json.NewDecoder(strings.NewReader(jsonStr)).Decode(&m)
	fmt.Println(m)

	// 这两种方式使用场景不一样，json.Unmarshal 和 json.NewDecoder 的区别主要是：
	// - json.Unmarshal 用于解析一个单独的 JSON 文档，而 json.NewDecoder 用于解析一个流式的 JSON 文档，例如从网络或文件中读取的数据。
	// - json.Unmarshal 需要先将 JSON 数据读取到内存中，然后再进行解析，而 json.NewDecoder 可以直接从 io.Reader 中读取并解析数据，不需要额外的内存分配。
	// - json.Unmarshal 只能解析一次，而 json.NewDecoder 可以解析多个连续的 JSON 文档。
	// 所以，一般来说，如果您的数据是从一个 io.Reader 流中获取的，或者您需要解析多个 JSON 文档，那么您应该使用 json.NewDecoder。如果您已经有了 JSON 数据在内存中，那么您可以使用 json.Unmarshal。

	// 说白了，如果我们是通过 API 进行交互，直接使用 resp.Body 的话，可以直接使用 json.NewDecoder(resp.Body).Decode(&cardGroups) 这种方式，而且不用再通过 io.Read 将 Body 转为 []byte。
}

func streaming() {
	// 示例字符串: {"Name":"Alice","Body":"Hello","Time":1294706395881547000}

	// 实例化一个流式解码器。将 JSON 数据映射到 Go 结构体中
	decoder := json.NewDecoder(os.Stdin)
	// 实例化一个流式编码器。将 Go 结构体中的数据映射为 JSON 数据
	encoder := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		if err := decoder.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := encoder.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	streaming()
	// UnmarshalAndDecode()
}
