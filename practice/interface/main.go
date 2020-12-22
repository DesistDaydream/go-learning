package main

import (
	"github.com/DesistDaydream/GoLearning/practice/interface/formula/rectangle"
	"github.com/DesistDaydream/GoLearning/practice/interface/formula/square"
	"github.com/DesistDaydream/GoLearning/practice/interface/salary/contract"
	"github.com/DesistDaydream/GoLearning/practice/interface/salary/permanent"
)

// Formula 计算不同图形的周长和面积
func Formula() {
	PrintResult("正方形", &square.Square{Side: 5})
	PrintResult("矩形", rectangle.Rectangle{Length: 4, Width: 3})
}

// Salary 计算工资薪酬
func Salary() {
	pemp1 := permanent.Permanent{EmpID: 1, Basicpay: 5000, Pf: 20}
	pemp2 := permanent.Permanent{EmpID: 2, Basicpay: 6000, Pf: 30}
	cemp1 := contract.Contract{EmpID: 3, Basicpay: 3000}
	// 接口是动态类型,本身的类型为nil，由于切片中的元素类型必须一样，所以接口切片中的元素也是接口并且类型为nil
	// 但是元素的类型会随着该元素的值的类型而变化，因为接口是动态类型，类型是可变的
	// 所以接口切片的底层切片类型依然是接口，但是表现出来的是已经改变过的类型
	// 使用接口的好处：如果不使用接口，每种类型的结构体就需要单独引用方法来计算总和，而没法放在一个数组中，进行迭代处理
	// 使用了接口后，可以把各种实现该接口的类型，放到该接口的数组里，进行统一处理。比如新加一个员工的类型就不用修改计算函数
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	TotalExpense(employees)
}

/*
使用接口的优势：`totalExpense`可以扩展新的员工类型，而不需要修改任何代码。
假如公司增加了一种新的员工类型`Freelancer`，它有着不同的薪资结构。
`Freelancer`只需传递到 totalExpense 的切片参数中，无需 totalExpense 函数本身进行修改。
只要 Freelancer 也实现了`SalaryCalculator`接口，`totalExpense`就能够实现其功能。
*/

// main 中的函数调用相当于模拟其他 package 调用本 package 的情景
func main() {
	// 计算图形面积和周长练习
	Formula()

	// 计算工资练习
	Salary()
}
