package main

import (
	"fmt"
)

// 关于使用 for 循环的一些注意事项
// https://zhuanlan.zhihu.com/p/345389953

// 使用循环迭代变量的指针
// 错误的做法
func forAttentionError() {
	in := []int{1, 2, 3}

	var out []*int
	for _, value := range in {
		out = append(out, &value)
	}

	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])

	// 这里的输出结果是:
	// Values: 3 3 3
	// Addresses: 0xc0000182b8 0xc0000182b8 0xc0000182b8
	// 这是因为 Go 语言中的 for 循环中，相当于声明了一个名为 value 的变量，每次迭代时，都是将元素的值改变，而不改变内存地址，所以在引用指针时，指向的都是同一个内存地址
	// 这就导致迭代完成后，该指针空间中存储的就是最后一个元素的值。
	// 想要解决这个问题并引用每次迭代的值的指针，可以参考 forPointerCorrect() 中的做法
}

// 正确的做法
func forPointerCorrect() {
	in := []int{1, 2, 3}

	var out []*int
	for _, value := range in {
		// 为解决 forAttentionError() 中的问题，将每次迭代都赋予一个新的变量，将这个新的变量的指针 appen 到切片中，这样，切片中所有元素的值，就都不一样了
		v := value
		out = append(out, &v)
	}

	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}

// 逃逸问题导致的指针地址变化
func forAttentionEscapeError() {
	in1 := []int{1, 2, 3}
	in2 := []string{"4", "5"}

	for i, v1 := range in1 {
		fmt.Printf("循环 %v\n", i+1)

		for _, v2 := range in2 {
			println("v1-Hex: ", &v1, "v2-Hex: ", &v2)
			// TODO: fmt 有逃逸问题？ https://juejin.cn/post/6955453411969990670，append 好像也有类似的内存逃逸现象
			// 正常情况下，v2 的指针应该也是一样的，但是用了 fmt 之后，指针的值在外层循环的下一次迭代中产生了变化
			fmt.Printf("v1-Dec: %v, v2-Dec: %v\n", &v1, &v2)
			// 从下面的输出结果可以看出来，注释掉 fmt 后，就没有逃逸现象，v2 指针的值是符合 for 的基础逻辑的: i.e. for 的 value 指针是固定的
			//
			// 输出结果如下，每次外层的迭代，都会改变 v2 指针的值：
			// 循环 1
			// v1-Hex:  0xc0001222c0 v2-Hex:  0xc00010a6a0
			// v1-Dec: 0xc0001222c0, v2-Dec: 0xc00010a6a0
			// v1-Hex:  0xc0001222c0 v2-Hex:  0xc00010a6a0
			// v1-Dec: 0xc0001222c0, v2-Dec: 0xc00010a6a0
			// 循环 2
			// v1-Hex:  0xc0001222c0 v2-Hex:  0xc00010a6b0
			// v1-Dec: 0xc0001222c0, v2-Dec: 0xc00010a6b0
			// v1-Hex:  0xc0001222c0 v2-Hex:  0xc00010a6b0
			// v1-Dec: 0xc0001222c0, v2-Dec: 0xc00010a6b0
			// 循环 3
			// v1-Hex:  0xc0001222c0 v2-Hex:  0xc00010a6c0
			// v1-Dec: 0xc0001222c0, v2-Dec: 0xc00010a6c0
			// v1-Hex:  0xc0001222c0 v2-Hex:  0xc00010a6c0
			// v1-Dec: 0xc0001222c0, v2-Dec: 0xc00010a6c0
			//
			// 注释掉 fmt 这行代码，则输出如下，v2 指针的值并不会改变：
			// 循环 1
			// v1-Hex:  0xc00004e6d8 v2-Hex:  0xc00004e710
			// v1-Hex:  0xc00004e6d8 v2-Hex:  0xc00004e710
			// 循环 2
			// v1-Hex:  0xc00004e6d8 v2-Hex:  0xc00004e710
			// v1-Hex:  0xc00004e6d8 v2-Hex:  0xc00004e710
			// 循环 3
			// v1-Hex:  0xc00004e6d8 v2-Hex:  0xc00004e710
			// v1-Hex:  0xc00004e6d8 v2-Hex:  0xc00004e710
		}
	}
}

type personOne struct {
	name string
	age  int
}

type personTwo struct {
	name string
	age  int
}

type company struct {
	ps []cp
}

type cp struct {
	personOne *personOne
	personTwo *personTwo
}

// for 与 结构体的测试，与基本类型逻辑一致
func forAttentionStructError() {
	ps1 := []personOne{
		{
			name: "zhangsan",
			age:  3,
		},
		{
			name: "lisi",
			age:  4,
		},
	}

	// var out company
	for _, v1 := range ps1 {
		ps2 := []personTwo{
			{
				name: "wangwu",
				age:  52,
			},
			{
				name: "zhaoliu",
				age:  62,
			},
		}
		for _, v2 := range ps2 {
			println("v1-Hex: ", &v1, "v2-Hex: ", &v2)
			// fmt.Printf("v1-Dec: %d, v2-Dec: %d\n", &v1, &v2)
		}
		// out.ps = append(out.ps, cp{
		// 	person: &p,
		// })
	}

	// fmt.Println(ps1)

	// for _, v := range out.ps {
	// 	fmt.Println(v.person)
	// }
	// fmt.Println("Values:", out.ps[0], out.ps[1])

	// TODO: ？？
	// 注意：若是多层 for 循环，那内层的循环则由于其他原因导致每次迭代的值的指针是不同的，就像下面这段代码
	// 这是因为，for 循环中的变量地址是固定的，但是在嵌套循环中，内部循环的变量地址会在每次外部循环迭代时重新分配。
	// 因此，在您提供的代码中，card 变量的地址在整个循环过程中保持不变，而 p 变量的地址则会在每次外部循环迭代时改变。
}
