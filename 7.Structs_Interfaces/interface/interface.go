package main

import (
	"fmt"
	"reflect"
)

// VowelsFinder 定义了一个名为VowelsFinder的接口
type VowelsFinder interface {
	FindVowels() []rune
}

// 定义了一个基础类型为字符串的名为MyString的类型
type (
	// MyString is
	MyString string
	// MyString1 is
	MyString1 string
)

// FindVowels MyString实现了接口VowelsFinder，因为MyString实现了FindVowels方法，且VowelsFinder接口中的所有方法该类型都实现
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

// FindVowels is
func (ms MyString1) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	name := MyString("Sam Anderson")
	name1 := MyString1("Tom Anderson")
	var v VowelsFinder
	// 接口已声明未使用的时候，接口变量的类型为nil(空),值也为nil(空)，接口的类型会随着值的不同而又不同的类型
	fmt.Println(reflect.TypeOf(name), reflect.TypeOf(v), v)
	fmt.Println(reflect.TypeOf(name1), reflect.TypeOf(v), v)
	// 由于MyString类型实现了VowelsFinder接口,所以该类型的变量与接口的变量可以互相赋值
	// 且接口变量的类型变为实现该接口的自定义类型
	v = name
	fmt.Printf("自定义类型为:%v,接口类型为:%v\n", reflect.TypeOf(name), reflect.TypeOf(v))
	v = name1
	fmt.Printf("自定义类型为:%v,接口类型为:%v\n", reflect.TypeOf(name), reflect.TypeOf(v))
	// 调用方法，把a、e、i、o、u在两个自定义类型变量的值所存在的字符的打印出来
	fmt.Printf("Vowels are %c", v.FindVowels())
}
