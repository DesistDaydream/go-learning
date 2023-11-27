package main

import (
	"fmt"
)

// 原始类型的结构体
type nativeStructString struct {
	Data []string `json:"data"`
}
type nativeStructInt struct {
	Data []int `json:"data"`
}

// genericStruct 是一个泛型分结构体
type genericStruct[T any] struct {
	Data []T
}

// newGenericStruct 创建一个泛型分页实例
func newGenericStruct[T any]() *genericStruct[T] {
	return &genericStruct[T]{
		// ！！！重点！！！这里为 nativeStruct.Data 确定了具体的类型。T 是什么类型，Data 就是什么类型。
		Data: make([]T, 0),
	}
}

// genGenericStructData 模拟生成数据并填充到 s.Data 中
func (s *genericStruct[T]) genGenericStructData() error {
	// 填充示例数据，主要是填充 s.Data
	// TODO: 示例数据好像并不好生成。因为需要对 T 具体是什么类型进行判断，可能需要反射？比如 gorm 的 Find() 等函数、json 中的相关编码解码函数。

	return nil
}

func Run() {
	// 实例化 GenericStruct 时传递了类型参数，让 genericStruct 的 T 变为 string 类型
	// newGenericStruct 内部使用了 Data: make([]T, 0)
	// 也就是说 genericStruct.Data 的类型实际变为了 []string。专业一点的说法是：T 被约束为 string 类型。
	genericStruct := newGenericStruct[string]()
	err := genericStruct.genGenericStructData()
	if err != nil {
		fmt.Println("生成数据错误:", err)
		return
	}

	// 由于 genericStruct.Data 的类型此时为 []string
	// 所以可以把 genericStruct.Data 赋值给 nativeStruct.Data，这俩都是 []string 类型
	stringData := nativeStructString{
		Data: genericStruct.Data, // 不用断言可以直接赋值。
	}

	fmt.Println(stringData)

	// 同样的，假如换了一个结构体，可以在 newGenericStruct 的时候改变类型参数
	genericStructInt := newGenericStruct[int]()
	err = genericStructInt.genGenericStructData()
	if err != nil {
		fmt.Println("生成数据错误:", err)
		return
	}
	intData := nativeStructInt{
		Data: genericStructInt.Data, // 不用断言可以直接赋值。
	}
	fmt.Println(intData)

	// 这样，就保证了有类似情况下，不用重复写 GenData 中的逻辑。也便于扩展和代码维护
}
