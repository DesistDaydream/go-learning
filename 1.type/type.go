package main

import (
	"fmt"
	"reflect"
)

func main() {
	var stringType string = "string"
	// 类型检测
	fmt.Println(reflect.TypeOf(stringType))
	fmt.Println(reflect.ValueOf(stringType))
}
