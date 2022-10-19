package main

import "fmt"

// 非泛型
type StructOne struct {
	Name string
}

type StructTwo struct {
	Name string
}

func GenDataOne(s *StructOne) string {
	return s.Name
}

func GenDataTwo(s *StructTwo) string {
	return s.Name
}

// TODO: 如何统一 GenDataOne 与 GenDataTwo？
type StructGeneric interface {
	StructOne | StructTwo
}

func GenData[T StructGeneric](s T) string {
	// 如何获取结构体中的 Name 属性？
	// return s.Name
	return ""
}

func Run() {
	var sOne StructOne
	var sTwo StructTwo
	GenDataOne(&sOne)
	GenDataTwo(&sTwo)

	// TODO: 泛型咋用呢？~
	fmt.Println("泛型输出：", GenData(sOne))
}
