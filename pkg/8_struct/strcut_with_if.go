package main

// 结构体中的字段是函数类型
type StructWithFunction struct {
	// 结构体中的字段是函数类型
	// 该函数类型的参数是一个int类型的参数，返回值是一个int类型的值
	// 该函数类型的名称是 FuncField
	FuncField func(int) int

	Name string
}

func NewStructWithFunction() *StructWithFunction {
	return &StructWithFunction{
		FuncField: func(i int) int {
			return i
		},
		Name: "DesistDaydream",
	}
}
