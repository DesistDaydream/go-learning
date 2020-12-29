package salary

// Permanent 永久员工类型，属性有ID和基本工资与附加工资
type Permanent struct {
	EmpID    int
	Basicpay int
	Pf       int
}

// CalculateSalary 永久员工的工资是基本工资和附加工资的综合
func (p Permanent) CalculateSalary() int {
	return p.Basicpay + p.Pf
}
