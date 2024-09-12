package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Code string `json:"code"`
	// ！！！重点！！！
	// 利用 json.RawMessage 类型来接收数据，这样就可以动态处理数据了
	Data json.RawMessage `json:"data"`
}

type SuccessData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ErrorData struct {
	Message string `json:"message"`
}

func main() {
	// 有时候我们需要动态得处理 struct 中的数据，比如 struct 的某个属性的类型可能有很多种
	jsonStr := `{"code":"success","data":{"name":"DesistDaydream","age":20}}`
	// jsonStr := `{"code":"error","data":{"message":"errorDemo"}}`

	var m Message
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		log.Fatalln(err)
	}

	switch m.Code {
	case "success":
		var data SuccessData
		err := json.Unmarshal(m.Data, &data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(data)
	case "error":
		var data ErrorData
		err := json.Unmarshal(m.Data, &data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(data)
	}
}
