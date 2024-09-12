package main

import (
	"fmt"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

// 类型断言
// 类型断言用于从接口类型的变量，获取该变量实际的类型（也可以说是检查该变量的类型）
// 类型断言的语法为：value, ok := element.(T)
// 其中，value 是转换后的值，ok 是一个布尔值，表示是否转换成功；element 是接口类型的变量，T 是目标类型
// 这个语法的意思是，如果 element 的类型是 T (e.g. string、int、etc.)，那么就返回 element 的值且返回值的类型为 T，也就是说 类型断言的语法为：value 的类型是 T
func typeAssertion(areaIntf interface{}) {
	// areaIntf 是 Square 类型吗？
	// 由于 areaIntf 是一个 interface 类型，所以现在还不知道areaIntf 此时的具体类型
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("areaIntf 的类型是: %T\n", t)
	}
	// areaIntf 是 Circle 类型吗？
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("areaIntf 的类型是: %T\n", u)
	} else {
		fmt.Println("areaIntf 不是 Circle 类型")
	}

	// 类型开关
	// 使用 switch 进行判断，type 为关键字，判断接口变量的类型
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Square 类型为: %T; 值为: %v\n", t, t)
	case *Circle:
		fmt.Printf("Circle 类型为: %T; 值为: %v\n", t, t)
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	typeAssertion(sq1)
}
