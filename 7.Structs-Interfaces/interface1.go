package main

import (
	"fmt"
	"reflect"
)

// 薪酬计算器接口，包含一个计算薪酬的方法
type SalaryCalculator interface {
	CalculateSalary() int
}

// 永久员工类型，属性有ID和基本工资与附加工资
type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

// 临时工类型，属性有ID和基本工资
type Contract struct {
	empId    int
	basicpay int
}

//永久员工的工资是基本工资和附加工资的综合
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//临时员工的工资只有基本工资
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

// 通过迭代SalaryCalculator切片并总结各个员工的工资来计算总费用
func totalExpense(s []SalaryCalculator) {
	expense := 0
	// 通过接口的切片s来获取其内每个元素的值v，根据其所对应的结构体类型，来引用相应的方法。
	// 虽然变量v的类型会变成不同的结构体类型，但是本质上，变量v依然是接口
	for index, v := range s {
		fmt.Printf("当前循环:%v，接口的类型转变为:%v\n",index,reflect.TypeOf(v))
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{empId: 1, basicpay: 5000, pf: 20}
	pemp2 := Permanent{empId: 2, basicpay: 6000, pf: 30}
	cemp1 := Contract{empId: 3, basicpay: 3000}
	// 接口是动态类型,本身的类型为nil，由于切片中的元素类型必须一样，所以接口切片中的元素也是接口并且类型为nil
	// 但是元素的类型会随着该元素的值的类型而变化，因为接口是动态类型，类型是可变的
	// 所以接口切片的底层切片类型依然是接口，但是表现出来的是已经改变过的类型
	// 使用接口的好处：如果不使用接口，每种类型的结构体就需要单独引用方法来计算总和，而没法放在一个数组中，进行迭代处理
	// 使用了接口后，可以把各种实现该接口的类型，放到该接口的数组里，进行统一处理。比如新加一个员工的类型就不用修改计算函数
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}

/*
使用接口的优势：`totalExpense`可以扩展新的员工类型，而不需要修改任何代码。
假如公司增加了一种新的员工类型`Freelancer`，它有着不同的薪资结构。
`Freelancer`只需传递到 totalExpense 的切片参数中，无需 totalExpense 函数本身进行修改。
只要 Freelancer 也实现了`SalaryCalculator`接口，`totalExpense`就能够实现其功能。
*/
