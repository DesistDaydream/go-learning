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
	// 仅声明了一个 *StructA 类型的变量
	var s1 *StructA
	fmt.Printf("s1 的内存地址：%p\n", s1) // 这里会输出 0x0，实际上代表的就是 <nil>
	fmt.Println(s1)
	// 这种方式 s1 为 nil，无法获取结构体中字段的值、也无法为字段赋值，将会报错：runtime error: invalid memory address or nil pointer dereference
	// fmt.Println(s1.FieldA)
	s1.ShowPtr()

	print("========\n")

	// 声明并初始化了一个 StructA 类型的变量
	s2 := StructA{}
	fmt.Printf("s2 的内存地址：%p\n", &s2)
	fmt.Println(&s2) // 这里将会输出 &{}，因为 s2 已被初始化，只不过字段值都为空(不是真的什么都没有，空也是有值的一种)。内存地址的表现形式被格式化了
	fmt.Println(s2.FieldA)
	s2.ShowPtr()

	print("========\n")

	// 在函数中，为 *StructA 类型的变量赋予了一个值
	s3 := NewStructA()
	fmt.Printf("s3 的内存地址：%p\n", s3)
	fmt.Println(s3.FieldA)
	s3.ShowPtr()

	print("========\n")
}
