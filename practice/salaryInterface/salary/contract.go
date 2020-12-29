package salary

// Contract 临时工类型，属性有ID和基本工资
type Contract struct {
	EmpID    int
	Basicpay int
}

// CalculateSalary 临时员工的工资只有基本工资
func (c Contract) CalculateSalary() int {
	return c.Basicpay
}
