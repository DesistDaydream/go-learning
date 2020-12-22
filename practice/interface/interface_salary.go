package main

import (
	"fmt"
	"reflect"
)

// SalaryCalculator 薪酬计算器接口，包含一个计算薪酬的方法
type SalaryCalculator interface {
	CalculateSalary() int
}

// 在 salary/ 下定义了两个结构体及其方法，结构体实现了 SalaryCalculator 接口

// TotalExpense 通过迭代 SalaryCalculator 切片并总结各个员工的工资来计算总费用
func TotalExpense(s []SalaryCalculator) {
	expense := 0
	// 通过接口的切片s来获取其内每个元素的值v，根据其所对应的结构体类型，来引用相应的方法。
	// 虽然变量v的类型会变成不同的结构体类型，但是本质上，变量v依然是接口
	for index, v := range s {
		fmt.Printf("当前循环:%v，接口的类型转变为:%v\n", index, reflect.TypeOf(v))
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}
