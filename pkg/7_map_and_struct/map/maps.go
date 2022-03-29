package main

import (
	"fmt"
)

// map 的基本使用方式
func mapDemo() {
	// map 声明
	var mapType map[string]int
	// map 初始化
	mapType = make(map[string]int)
	// map 赋值
	mapType["key1"] = 10
	mapType["key2"] = 11
	fmt.Printf(" map 长度: %v\n map 第一个值: %v\n map 第二个值: %v\n", len(mapType), mapType["key1"], mapType["key2"])
	// 删除 mapType 这个 map 中 key 为 key1 的值
	delete(mapType, "key1")
	fmt.Printf("删除了 map 中的一个 k/v:\n 删除后 map 长度: %v\n 删除后 map 第一个值: %v\n 删除后 map 第二个值: %v\n", len(mapType), mapType["key1"], mapType["key2"])

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
	fmt.Println("遍历输出:")
	for key, val := range elements {
		fmt.Println(key, ":", val)
	}
}

func main() {
	fmt.Println("1.map 的基本使用方法：")
	mapDemo()
}
