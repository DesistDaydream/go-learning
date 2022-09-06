package main

import "fmt"

type StructA struct {
	FieldA string
}

func NewStructA() *StructA {
	return &StructA{}
}

func (s *StructA) ShowPtr() {
	fmt.Println("接收器的值取决于调用时的值：", s)
}

func StructAndPointer() {
	// 声明并初始化了一个 StructA 类型的变量
	s1 := StructA{}
	fmt.Printf("s1 的内存地址：%p\n", &s1) // 若使用 fmt.Println(&s1)，将会输出 &{}，因为 s1 已被初始化，只不过字段值都为空(不是真的什么都没有，空也是有值得一种)。所以这种方式可以获取到内存地址
	fmt.Println(&s1)
	fmt.Println(s1.FieldA)
	s1.ShowPtr()

	print("========\n")

	// 仅声明了一个 *StructA 类型的变量
	var s2 *StructA
	fmt.Printf("s2 的内存地址：%p\n", s2) // 这里会输出 0x0，实际上就是 <nil>
	// 这种方式 s2 为 nil，无法获取结构体中字段的值、也无法为字段赋值，将会报错：runtime error: invalid memory address or nil pointer dereference
	// fmt.Println(s2.FieldA)
	s2.ShowPtr()

	print("========\n")

	// 在函数中，为 *StructA 中赋予了一个值
	s3 := NewStructA()
	fmt.Printf("s3 的内存地址：%p\n", s3)
	fmt.Println(s3.FieldA)
	s3.ShowPtr()

	print("========\n")
}
