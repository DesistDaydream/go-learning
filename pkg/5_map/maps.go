package main

import (
	"encoding/json"
	"fmt"
)

// map 的基本使用方式
func mapDemo() {
	fmt.Println("1.map 的基本使用方法：")
	s1 := []int{5, 9, 12, 5, 24, 7, 8, 4, 17, 9, 2, 4, 13, 1, 24, 6, 2, 31, 12, 5, 10, 24}
	m1 := make(map[int]int)
	// var s2 []int
	// var max int
	// var s3 []int

	for _, v := range s1 {
		fmt.Println(m1)
		if m1[v] != 0 {
			m1[v]++
		} else {
			m1[v] = 1
		}
	}
	fmt.Println(m1)

	// for _, v := range m1 {
	// 	s2 = append(s2, v)
	// }
	// max = s2[0]
	// for i := 0; i < len(s2); i++ {
	// 	if max < s2[i] {
	// 		max = s2[i]
	// 	}
	// }
	// fmt.Println(max)

	// for k, v := range m1 {
	// 	if v == max {
	// 		s3 = append(s3, k)
	// 	}
	// }
	// fmt.Println(s3)
}

// 遍历 map
func ForRangeMap() {
	data := map[string]string{
		"username": "DesistDaydream",
		"password": "My@Password",
	}

	// 使用 for...range... 遍历 map 类型的数据时
	// for 后的 第一个变量是 map 中每个元素的 key，第二个变量是 map 中每个元素 key 所对应的 value
	for key, val := range data {
		fmt.Println(key, ":", val)
	}
}

// 将 map 类型数据转为 JSON 格式数据
func mapToJSON() {
	data := map[string]string{
		"username": "DesistDaydream",
		"password": "My@Password",
	}

	jsonData, _ := json.Marshal(data)

	fmt.Println(string(jsonData))

	// 输出如下内容：{"password":"My@Password","username":"DesistDaydream"}
}

// 评估 map 是否可以正常获取值
func evaluatesMap() {
	data := map[string]string{
		"username": "DesistDaydream",
		"password": "My@Password",
	}

	// 评估 map 是否可以正常获取值
	if val, ok := data["username"]; ok {
		fmt.Println("map 中 key 为 username 的值为：", val)
	} else {
		fmt.Println("map 中 key 为 username 的值不存在")
	}
}

func main() {
	// map 声明
	var mapType map[string]int
	// map 初始化，通常声明与初始化合并，这里仅作为演示
	mapType = make(map[string]int)
	// map 赋值
	mapType["key1"] = 10
	mapType["key2"] = 11
	// map 引用
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
}
