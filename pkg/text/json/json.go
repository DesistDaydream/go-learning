package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 将字符串格式的 JSON 数组转为真实的数组类型的数据
func StringToSlice() {
	jsonArrayStr := "[\"AAA\",\"BBB\"]"
	realSlice := []string{}
	err := json.Unmarshal([]byte(jsonArrayStr), &realSlice)
	if err != nil {
		log.Fatalln(err)
	}
	for _, slice := range realSlice {
		fmt.Println(slice)
	}
}

func main() {
	StringToSlice()
}
