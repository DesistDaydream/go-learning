package main

import (
	"fmt"
	"reflect"
)

type NeedReflectStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ReflectAndStruct() {
	s := &NeedReflectStruct{
		Name: "DesistDaydream",
		Age:  10,
	}
	tP := reflect.TypeOf(s)
	vP := reflect.ValueOf(s)
	fmt.Printf("类型：%v，值：%v\n", tP, vP)

	// 注意：指针类型比较特殊，需要通过 Elem() 方法以获取内部子元素的类型
	// Elem() 只能由 Array, Chan, Map, Pointer, Slice 这几种类型调用，若是其他类型，程序将会 Panic
	// 下面的代码将会 Panic，因为指针类型不是结构体类型，无法使用 Filed() 方法
	// fmt.Println(tP.Field(0).Name)
	// 正确的代码是应该先通过 Elem() 方法获取到指针类型内的子类型
	t := tP.Elem()
	v := vP.Elem()
	fmt.Printf("类型：%v，值：%v\n", t, v)
	// 拿到 类型实例 后，我们可以通过多种方法获取类型的信息
	// NumField() 方法只适用于 Struct 类型，返回结构体类型中所有字段的总数。
	for i := 0; i < t.NumField(); i++ {
		// Field() 方法只适用于 Struct 类型，返回结构体各个字段的信息
		fieldInfo := t.Field(i)
		// 字段信息包括该字段的 名称、包路径、数据类型、标签 等等
		fmt.Printf("字段名称：%v\n", fieldInfo.Name)
		fmt.Printf("字段包路径：%v\n", fieldInfo.PkgPath)
		fmt.Printf("字段标签：%v\n", fieldInfo.Tag)
	}

}

func Reflect() {
	var strVar string
	// 获取一个变量的类型
	t := reflect.TypeOf(strVar)
	// 获取一个变量的值
	v := reflect.ValueOf(strVar)
	fmt.Println(t, v)
}

func main() {
	// Reflect()
	ReflectAndStruct()
}
