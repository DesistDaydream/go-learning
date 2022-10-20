package main

import (
	"fmt"
	"reflect"
)

type NeedReflectStruct struct {
	Name string `json:"name" yaml:"name"`
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
		fieldInstance := t.Field(i)      // 对于字段来说，Field() 返回一个关于结构体字段的实例
		fieldValueInstance := v.Field(i) // 对于值来说，Field() 返回一个关于结构体字段的值的实例

		fmt.Printf("字段实例: %v\n", fieldInstance)
		fmt.Printf("字段值实例: %v\n", fieldValueInstance)

		// 字段信息包括该字段的 名称、包路径、数据类型、标签 等等
		fmt.Printf("第 %v 个字段名称：%v\n", i+1, fieldInstance.Name)
		fmt.Printf("第 %v 个字段标签：%v\n", i+1, fieldInstance.Tag)
		// 获取结构体中 Tag 的值
		fmt.Printf("第 %v 个字段标签中名为 json 的值: %v\n", i+1, fieldInstance.Tag.Get("json"))

		// 通过 值实例 的 Interface() 方法获取字段的值
		fmt.Printf("第 %v 个字段值：%v\n", i+1, fieldValueInstance.Interface())
	}

	// Kind() 方法用来获取目标的类型，常用来执行比较运算
	if t.Kind() == reflect.Struct {
		fmt.Println("这是一个 Struct 类型")
	}

	// SetXXX() 方法可以让我们修改变量的值
	// 这里我们通过 FieldByName() 找到名为 Name 的字段，然后使用 SetString() 修改 Name 字段的值
	v.FieldByName("Name").SetString("新名字")
	fmt.Println(s)

}

func Reflect() {
	var intVar int
	// 获取一个变量的类型。返回 Type{} 接口，该接口有很多方法用以获取该变量“类型”的相关信息。
	t := reflect.TypeOf(intVar)
	// 获取一个变量的值。返回 Value{} 结构体，该结构体也有很多方法用以获取该变量“值”的相关信息。
	v := reflect.ValueOf(intVar)

	fmt.Println("类型: ", t)
	fmt.Println("值: ", v)
}

func main() {
	Reflect()
	// ReflectAndStruct()
}
