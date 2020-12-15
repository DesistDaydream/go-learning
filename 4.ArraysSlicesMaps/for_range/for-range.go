package main

import "fmt"

func main() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	//For-Range结构应用于数组和切片，index为索引，season为该索引位置元素的值
	for index1, season := range seasons {
		fmt.Printf("Season %d is: %s\n", index1, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

	for index1 := range seasons {
		fmt.Printf("%d ", index1)
	}
}