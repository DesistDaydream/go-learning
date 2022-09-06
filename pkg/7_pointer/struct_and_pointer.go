package main

import "fmt"

type StructA struct {
	FieldA string
}

func NewStructA() *StructA {
	return &StructA{}
}

func (s *StructA) Add() {
	fmt.Println("接收器的值取决于调用时的值：", s)
}

func StructAndPointer() {
	// 声明并初始化了一个 StructA 类型的变量
	s1 := StructA{}
	fmt.Println(&s1)
	fmt.Println(&s1.FieldA)
	s1.Add()

	// 仅声明了一个 *StructA 类型的变量
	var s2 *StructA
	fmt.Println(s2)
	// 这种方式 s2 为 nil，无法获取结构体中字段的值，将会报错：runtime error: invalid memory address or nil pointer dereference
	// fmt.Println(s2.FieldA)
	s2.Add()

	// 在函数中，为 *StructA 中赋予了一个值
	s3 := NewStructA()
	fmt.Println(s3)
	fmt.Println(s3.FieldA)
}
