package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 切片是对数组的引用，使用`make`函数创建一个长度为3的string类型的切片，初始化为零值
	s := make([]string, 3)
	fmt.Println("emp:", s)
	// 可以和数组一样，设置、得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s, reflect.TypeOf(s[0]))
	fmt.Println("get:", s[2])
	// `len`函数返回slice的长度
	fmt.Println("len:", len(s))
	// append函数返回一个包含了一个或者多个新值的切片
	s = append(s, "d")
	fmt.Println("apd:", s)
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	// 切片可以被拷贝，创建一个空的和s有相同长度的切片c，并将s复制给c
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	// 使用`SliceID[LOW:HIGH]`语法对指定的切片进行`切片`操作
	// e.g.下面得到一个包含元素s[2],s[3],s[4]的切片
	l := s[2:5]
	fmt.Println("sl1:", l)
	// 切片l中的数据从s[0]到s[5](不包含s[5])
	l = s[:5]
	fmt.Println("sl2:", l)
	// 切片l中的数据从s[2]开始到最后(包含s[2])
	l = s[2:]
	fmt.Println("sl3:", l)
	//可以在一行代码中声明并初始化一个slice变量
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// Slice 可以组成多维数据结构。
	// 内部的 slice 长度可以不 一致，这和多维数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
