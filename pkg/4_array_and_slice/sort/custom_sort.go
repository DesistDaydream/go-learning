package main

import (
	"fmt"
	"sort"
)

// 想要实现自己的排序逻辑，我们需要让一个类型实现 sort.Interface{} 接口
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func CustomSort() {
	var fruits byLength = []string{"peach", "banana", "kiwi"}
	// 有时候，我们可能想根据自然顺序以外的方式来对集合进行排序。
	// 例如，假设我们要按字符串的长度而不是按字母顺序对它们进行排序。
	// 用白话说，就是根据自己的需求决定排序的逻辑
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}

type Person struct {
	Name string
	Age  int
}

type Persons []Person

func SimpleCustomSort() {
	// 有时候我们还需要对 Struct 进行排序，当代码中有大量的 Struct 的时候，如果每个 Struct 都要实现三个方法，是非常不友好的。
	// 且排序涉及到两个变量的值，而 Struct 可能包含多个属性，程序并不知道我们想以哪个属性或哪几个属性作为衡量大小的标准。
	// 此时我们就需要告诉程序比较逻辑
	// 在 sort 包中，有 Slice()、SliceStable()、SliceIsSorted()、Search() 这几个函数可以满足上述需求
	persons := Persons{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(persons, func(i, j int) bool {
		// 自己定义的排序逻辑
		// 这里是根据 Struct 中 Age 字段的值的大小进行升序排序
		return persons[i].Age < persons[j].Age
	})

	fmt.Println(persons)
}
