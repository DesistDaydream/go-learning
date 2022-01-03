package main

import (
	"fmt"
	"reflect"
)

// map 的基本使用方式
func mapsDemo() {
	// 第一种 map 的使用方法
	var x map[string]int
	fmt.Println(reflect.TypeOf(x))
	x = make(map[string]int)
	fmt.Println(reflect.TypeOf(x))
	x["key1"] = 10
	x["key2"] = 11
	fmt.Println(len(x), x["key1"], x["key2"])
	// 删除 x 这个 map 中 key 为 key1 的值
	delete(x, "key1")
	fmt.Println(len(x), x["key1"], x["key2"])

	// 第二种 map 的使用方法，可以省去使用 make 分配内存的过程
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Printf("map 的长度为：%v\nmap 中 key 为 H 的值为：%v\n", len(elements), elements["H"])
	for key, val := range elements {
		fmt.Println("遍历输出", key, ":", val)
	}
}

func main() {
	fmt.Println("1.map 的基本使用方法：")
	mapsDemo()
}
