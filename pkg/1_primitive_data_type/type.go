package main

import (
	"fmt"
	"reflect"
	"strconv"
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
func typeConversions() {
	// string 转 []byte
	str := "hello"
	bytes := []byte(str)

	// []byte 转 string
	str2 := string(bytes)

	fmt.Println(str2)

	// 指针转换，需要在类型两边添加小括号
	type stringValue string
	ptr := (*stringValue)(&str2)
	fmt.Println(ptr)
}

// 使用 strconv 标准库来实现字符串类型与其他基本数据类型之间的类型转换。
// strconv 全称是 string q
func stringIntFloat() {
	strInt := "1"
	strFloat := "1.234"

	// Atoi() 与 Itoa() 方法是最常用的转换
	// Atoi() 相当于 ParseInt(s, 10, 0)
	// Itoa() 相当于 FormatInt(int64(i), 10)
	// 将 string 类型转换为 int 类型
	int1, _ := strconv.Atoi(strInt)
	// 将 int 类型转换为string类型
	str3 := strconv.Itoa(int1)

	// string 转 float32 和 float64
	float32One, _ := strconv.ParseFloat(strFloat, 32)
	float64One, _ := strconv.ParseFloat(strFloat, 64)

	fmt.Println(str3)
	fmt.Println(float32One)
	fmt.Println(float64One)

	// 其他基础类型转字符串类型
	var int64Str int64 = 20
	// int64 转 string
	fmt.Println(strconv.FormatInt(int64Str, 10))
}

// 数据类型转换，字符串转Unicode，Unicode转字符串
func unicode() {
	sText := "断念梦"
	// 字符串转 Unicode
	textQuoted := strconv.QuoteToASCII(sText)
	fmt.Println(textQuoted)

	// Unicode 转字符串
	str, _ := strconv.Unquote(textQuoted)
	fmt.Println(str)
}
