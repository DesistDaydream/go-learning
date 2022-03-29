package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	// 值为22
	fmt.Printf("The sum is: %d\n", two1.AddThem())
	// 值为42
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	// 值为7
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}
