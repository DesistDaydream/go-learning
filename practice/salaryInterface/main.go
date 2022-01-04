package main

import (
	"github.com/DesistDaydream/go-learning/practice/salaryInterface/salary"
	"github.com/DesistDaydream/go-learning/practice/salaryInterface/salaryInterface"
)

func main() {
	// Salary 计算工资薪酬
	pemp1 := salary.Permanent{EmpID: 1, Basicpay: 5000, Pf: 20}
	pemp2 := salary.Permanent{EmpID: 2, Basicpay: 6000, Pf: 30}
	cemp1 := salary.Contract{EmpID: 3, Basicpay: 3000}
	// 接口是动态类型,本身的类型为nil，由于切片中的元素类型必须一样，所以接口切片中的元素也是接口并且类型为nil
	// 但是元素的类型会随着该元素的值的类型而变化，因为接口是动态类型，类型是可变的
	// 所以接口切片的底层切片类型依然是接口，但是表现出来的是已经改变过的类型
	// 使用接口的好处：如果不使用接口，每种类型的结构体就需要单独引用方法来计算总和，而没法放在一个数组中，进行迭代处理
	// 使用了接口后，可以把各种实现该接口的类型，放到该接口的数组里，进行统一处理。比如新加一个员工的类型就不用修改计算函数
	employees := []salaryInterface.SalaryCalculator{pemp1, pemp2, cemp1}
	salaryInterface.TotalExpense(employees)
}

/*
使用接口的优势：`totalExpense`可以扩展新的员工类型，而不需要修改任何代码。
假如公司增加了一种新的员工类型`Freelancer`，它有着不同的薪资结构。
`Freelancer`只需传递到 totalExpense 的切片参数中，无需 totalExpense 函数本身进行修改。
只要 Freelancer 也实现了`SalaryCalculator`接口，`totalExpense`就能够实现其功能。
*/
