package main

import "fmt"

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

type person struct {
	name string
	age  int
}

type company struct {
	ps []cp
}

type cp struct {
	person *person
}

// 使用循环迭代变量的指针
// 错误的做法
func forAttentionStructError() {
	ps := []person{
		{
			name: "lch",
			age:  30,
		},
		{
			name: "zn",
			age:  300,
		},
	}

	var out company
	for _, p := range ps {
		println(&p)
		fmt.Println(&p)
		out.ps = append(out.ps, cp{
			person: &p,
		})
	}

	fmt.Println(ps)

	for _, v := range out.ps {
		fmt.Println(v.person)
	}
	fmt.Println("Values:", out.ps[0], out.ps[1])

	// 注意：若是多层 for 循环，那内层的循环则由于其他原因导致每次迭代的值的指针是不同的，就像下面这段代码
	// for _, card := range cards.Data {
	// 	for _, p := range products.Data {
	// 		println(&card)
	// 		println(&p)
	// 	}
	// }
	// 输出的结果是：
	// 0xc0003a2d80
	// 0xc00028a500
	// 0xc0003a2d80
	// 0xc000168140
	// 0xc0003a2d80
	// 0xc00028a640
	// 0xc0003a2d80
	// 0xc000168280
	// 这是因为，for 循环中的变量地址是固定的，但是在嵌套循环中，内部循环的变量地址会在每次外部循环迭代时重新分配。
	// 因此，在您提供的代码中，card 变量的地址在整个循环过程中保持不变，而 p 变量的地址则会在每次外部循环迭代时改变。
}
