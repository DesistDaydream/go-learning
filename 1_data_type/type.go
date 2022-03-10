package main

import (
	"fmt"
	"reflect"
)

// 基本数据类型
var intType = 12345
var stringType = "string"
var booleanType = true

// 复合数据类型
var sliceType []string
var mapType = make(map[string]string)

var structType struct{}
var interfaceType interface{}

var channelType = make(chan string)

var pointType = new(int)

func funcType() {}

// 数据类型检查
func dataType() {
	// 基础数据类型检测
	fmt.Println("检测数值变量的数据类型：", reflect.TypeOf(intType))
	fmt.Println("检测字符串变量的数据类型：", reflect.TypeOf(stringType))
	fmt.Println("检测布尔变量的数据类型：", reflect.TypeOf(booleanType))

	// 复合数据类型检测
	fmt.Println("检测数组变量的数据类型：")
	fmt.Println("检测切片变量的数据类型：", reflect.TypeOf(sliceType))
	fmt.Println("检测字典变量的数据类型：", reflect.TypeOf(mapType))
	fmt.Println("检测函数变量的数据类型：", reflect.TypeOf(funcType))
	fmt.Println("检测结构体变量的数据类型：", reflect.TypeOf(structType))
	fmt.Println("检测接口变量的数据类型：", reflect.TypeOf(interfaceType))
	fmt.Println("检测通道变量的数据类型：", reflect.TypeOf(channelType))
	fmt.Println("检测指针变量的数据类型：", reflect.TypeOf(pointType))
}

// 数据类型转换
func typeTransformation() {
	// string to []byte
	s1 := "hello"
	b := []byte(s1)

	// []byte to string
	s2 := string(b)

	fmt.Println(s2)
}

func main() {
	// 数据类型检查
	dataType()

	// 数据类型转换
	typeTransformation()
}
