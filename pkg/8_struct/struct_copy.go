package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type User struct {
	Name        string
	Role        string
	Age         int32
	EmployeCode int64 `copier:"EmployeNum"` // 指定字段名。

	// Explicitly ignored in the destination struct.
	Salary int
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

// Tags in the destination Struct provide instructions to copier.Copy to ignore
// or enforce copying and to panic or return an error if a field was not copied.
type Employee struct {
	// 这个 Tag 告诉 copier，如果字段没有被复制，则 panic
	Name string `copier:"must"`

	// 这个 Tag 告诉 copier，如果字段没有被复制，则返回一个错误
	Age int32 `copier:"must,nopanic"`

	// 这个 Tag 告诉 copier，忽略该字段，即复制时不复制该字段的值
	Salary int `copier:"-"`

	DoubleAge int32
	EmployeId int64 `copier:"EmployeNum"` // 指定字段名。
	SuperRole string
}

func (employee *Employee) Role(role string) {
	employee.SuperRole = "Super " + role
}

func StructCopy() {
	// 我们经常会遇上需要将一个结构体中的值拷贝到另一个相似的结构体中的情况
	// 此时可以使用一个第三方库来简化操作
	var (
		user      = User{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 200000}
		users     = []User{{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 100000}, {Name: "jinzhu 2", Age: 30, Role: "Dev", Salary: 60000}}
		employee  = Employee{Salary: 150000}
		employees = []Employee{}
	)

	// 一、struct 拷贝到 struct
	// 将第二个参数的值拷贝到第一个参数中
	copier.Copy(&employee, &user)

	fmt.Printf("%#v \n", employee)
	// Employee{
	//    Name: "Jinzhu",           // Copy from field
	//    Age: 18,                  // Copy from field
	//    Salary:150000,            // Copying explicitly ignored
	//    DoubleAge: 36,            // Copy from method
	//    EmployeeId: 0,            // Ignored
	//    SuperRole: "Super Admin", // Copy to method
	// }

	// 二、struct 拷贝到 slice
	copier.Copy(&employees, &user)

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeId: 0, SuperRole: "Super Admin"}
	// }

	// 三、slice 拷贝到 slice
	employees = []Employee{}
	copier.Copy(&employees, &users)

	fmt.Printf("%#v \n", employees)
	// []Employee{
	//   {Name: "Jinzhu", Age: 18, Salary:0, DoubleAge: 36, EmployeId: 0, SuperRole: "Super Admin"},
	//   {Name: "jinzhu 2", Age: 30, Salary:0, DoubleAge: 60, EmployeId: 0, SuperRole: "Super Dev"},
	// }

	// 四、map 拷贝到 map
	map1 := map[int]int{3: 6, 4: 8}
	map2 := map[int32]int8{}
	copier.Copy(&map2, map1)

	fmt.Printf("%#v \n", map2)
	// map[int32]int8{3:6, 4:8}
}
