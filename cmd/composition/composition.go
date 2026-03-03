// 组合的实践
package main

import "fmt"

// 当我们有多个结构体，且这些结构体有一些共同的字段时，我们可以使用组合来避免重复定义字段
// 假如，我们有两个安全数据
type MySecurityDataInfo struct {
	Version    string
	Code       string
	Name       string
	MyProperty string
}

func (m MySecurityDataInfo) GetCode() string {
	return m.Code
}

type YourSecurityDataInfo struct {
	Version      string
	Code         string
	Name         string
	YourProperty string
}

func (y YourSecurityDataInfo) GetCode() string {
	return y.Code
}

type SecurityDataInfo interface {
	GetCode() string
}

func Demo() {
	demo := []SecurityDataInfo{
		MySecurityDataInfo{
			Version:    "1.0",
			Code:       "123456",
			Name:       "DesistDaydream",
			MyProperty: "DesistDaydreamProperty",
		},
		YourSecurityDataInfo{
			Version:      "1.0",
			Code:         "654321",
			Name:         "Your",
			YourProperty: "YourProperty",
		},
	}
	// 当我想要在某个循环中，都获取到 Code 属性的值
	// 我们首先想到的是要设计一个接口（type SecurityDataInfo interface）来统一实现两个结构体
	// 然后，我们分别实现这两个结构体的 GetCode 方法
	// 最后，我们可以在循环中，通过接口调用 GetCode 方法来获取 Code 字段
	for _, v := range demo {
		fmt.Println(v.GetCode())
	}

	// 随着开发，我们发现还想都获取到 Name 属性的值
	// 然后还要再添加一个 GetName 方法。
	// 每多一个属性，我们就需要添加一个对应的方法。
	// 那么，有没有什么办法可以避免这个问题呢？
	// **组合** 可以解决这个问题。详见下面的代码
}

// ########
// 组合的实践 - 提取共同字段
// ########

// 我们将两个 struct 中的共同的属性提取出来，定义到一个新的结构体中
// 然后，我们在两个 struct 中嵌入这个新的 struct
// 这样，我们就可以在两个 struct 中都访问到共同的字段了
type BaseSecurityDataInfo struct {
	Version string
	Code    string
	Name    string
}

func (m BaseSecurityDataInfo) GetBaseInfo() BaseSecurityDataInfo {
	return m
}

// 通过**组合**的方式，让 MySecurityDataInfoComposition 与 BaseSecurityDataInfo 组合起来
// 这样，我们就可以在 MySecurityDataInfoComposition 中访问到 BaseSecurityDataInfo 中的字段了
// 同时，还可以使用 BaseSecurityDataInfo 下的 GetBaseInfo 方法来获取 BaseSecurityDataInfo 中的字段
type MySecurityDataInfoComposition struct {
	BaseSecurityDataInfo
	MyProperty string
}

// 也可以在每个 struct 中实现 GetBaseInfo 方法
// TODO: 思考一下，这两种方式有什么区别？
// func (m MySecurityDataInfoComposition) GetBaseInfo() BaseSecurityDataInfo {
// 	return m.BaseSecurityDataInfo
// }

type YourSecurityDataInfoComposition struct {
	BaseSecurityDataInfo
	YourProperty string
}

// func (y YourSecurityDataInfoComposition) GetBaseInfo() BaseSecurityDataInfo {
// 	return y.BaseSecurityDataInfo
// }

type SecurityDataInfoComposition interface {
	GetBaseInfo() BaseSecurityDataInfo
}

func CompositionDemo() {
	demo := []SecurityDataInfoComposition{
		MySecurityDataInfoComposition{
			BaseSecurityDataInfo: BaseSecurityDataInfo{
				Version: "1.0",
				Code:    "123456",
				Name:    "DesistDaydream",
			},
			MyProperty: "DesistDaydreamProperty",
		},
		YourSecurityDataInfoComposition{
			BaseSecurityDataInfo: BaseSecurityDataInfo{
				Version: "1.0",
				Code:    "654321",
				Name:    "Your",
			},
			YourProperty: "YourProperty",
		},
	}

	for _, v := range demo {
		fmt.Println(v.GetBaseInfo().Version)
		fmt.Println(v.GetBaseInfo().Code)
		fmt.Println(v.GetBaseInfo().Name)
	}
}
