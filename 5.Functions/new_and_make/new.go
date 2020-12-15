package main

import "fmt"

// Person 定义人的属性
type Person struct {
	name string
	age  int
}

// Values 为 Person 结构体复制
func (p *Person) Values() {
	p.name = "DesistDaydream"
	p.age = 18
}

func main() {
	// 返回一个 Person 结构体的指针
	p := new(Person)
	//
	p.Values()
	fmt.Printf("姓名: %v\n年龄：%v\n", p.name, p.age)
}
