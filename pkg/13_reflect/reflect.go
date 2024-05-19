package main

import (
	"fmt"
	"log"
	"reflect"
)

// Reflect(反射) 功能基本示例
func ReflectDemo() {
	var intVar int
	// Go 语言中 Reflect 功能的核心依赖两个抽象:
	// - reflect.Type 接口 # reflect.Type 是 Go 语言中所有类型的抽象，比如 int、string、struct、etc.
	// - reflect.Value 结构体 # reflect.Value 是 Go 语言中某个值的抽象，可以通过 reflect.Value.Type() 获取该值对应的 reflect.Type 接口。
	//
	// 由于 Type 和 Value 的所有方法并非适用于所有类型，在调用方法前推荐先使用 t.Kind 找出 Type 或 Value 所属的 Kind。
	// Type 与 Kind 的区别可以参考下面的 ReflectKindAndType() 示例

	// 获取一个变量的类型。返回 reflect.Type{} 接口，该接口有很多方法用以获取该变量“类型”的相关信息。
	t := reflect.TypeOf(intVar)
	// 获取一个变量的值。返回 reflect.Value{} 结构体，该结构体也有很多方法用以获取该变量“值”的相关信息。
	v := reflect.ValueOf(intVar)
	// 注意：t 与 v 下的方法并不适用于所有类型，比如 Field() 方法只作用于 Struct 类型，如果变量类型不是 Struct，调用该方法时将会 Panic
	// ！！！但是由于反射的机制，在我们编写代码时，是无法知道将要调用方法的实例是什么类型，只有运行起来之后才会知道，这点要万分注意！！！
	fmt.Println("类型: ", t)
	fmt.Println("值: ", v)

	// 如果执行如下代码，将会报错：panic: reflect: Field of non-struct type int
	// t.Field(0)
}

// reflect 可以获取变量的类型，一个变量的类型分为两种: Kind 与 Type。Type 是 Kind 的子集
// 类似于 公交车（Type） 和 交通工具（Kind） 的关系。
// 比如一个定义了一个 structDemo 的 struct，这个 structDemo 就是 Type；struct 就是 Kind。
func ReflectKindAndType() {
	type structType struct {
	}
	var structVar structType
	v := reflect.ValueOf(structVar)
	log.Printf("Kind 类型: %v", v.Kind())
	log.Printf("Type 类型: %v", v.Type())
}

func main() {
	ReflectDemo()
	ReflectKindAndType()
}
