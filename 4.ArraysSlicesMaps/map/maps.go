package main

import (
	"fmt"
)

func main() {
	fmt.Println("1.map的基本使用方法：")
	maps1()
}

// maps的基本使用方式
func maps1() {
	// 第一种map的使用方法
	var x map[string]int
	x = make(map[string]int)
	x["key"] = 10
	x["key1"] = 11
	fmt.Println(len(x), x["key"], x["key1"])
	delete(x, "key")
	fmt.Println(len(x), x["key"], x["key1"])
	
	// 第二种名片map的使用方法
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
	fmt.Println(len(elements), elements["H"])	
	for key, val := range elements {
		fmt.Println("遍历输出",key, ":",val)
	}
}
