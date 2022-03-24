package main

import "fmt"

// Person 定义人的属性
type Person struct {
	name string
	age  int
}

// Values 为 Person 结构体赋值
func (p *Person) Values() {
	p.name = "DesistDaydream"
	p.age = 18
}

func main() {
	// 返回一个 Person 结构体的指针，以便后面的方法可以直接操纵该结构的属性的值
	// 这种行为一般成为实例化，将 Person 实例、返回一个 Person 实例、等等
	p := new(Person)
	// 调用方法来处理 Person 属性的值
	p.Values()
	// 打印结构体属性中的值
	fmt.Printf("姓名: %v\n年龄：%v\n", p.name, p.age)
}
