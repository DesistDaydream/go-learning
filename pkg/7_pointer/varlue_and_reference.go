package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func Call() {
	// Go 通过指针，允许程序通过`引用传递`来传递值和数据结构
	i := 1
	fmt.Println("initial:", i)

	// 通过值传递，在函数内修改不会影响原始值
	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过引用传递，在函数内修改会影响原始值
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 注意一下，虽然从逻辑上看，Go 确实做了引用传递，但是实际上，Go 是将内存地址复制了一份，然后传到函数中
	// 所以，虽然逻辑上是引用传递，但是实际上是传递了一个值，只不过这个值是内存地址。
}

// 逻辑上的引用传递
func setNameV1(user *User) {
	fmt.Printf("v1: %p\n", user)    // 0xc0000a4018  与 init的地址一致
	fmt.Printf("v1_p: %p\n", &user) // 0xc0000ac020，这里其实是指的值为 0xc0000a4018 的 user 变量的指针地址
	user.Name = "test_v1"
}

// 值传递
func setNameV2(user User) {
	fmt.Printf("v2_p: %p\n", &user) //0xc0000a4030
	user.Name = "test_v2"
}

func CallTwo() {
	u := User{Name: "init"}

	uPointer := &u
	fmt.Printf("init: %p \n", uPointer) //0xc0000a4018
	setNameV1(uPointer)
	setNameV2(u)
}
