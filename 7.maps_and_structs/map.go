package main

import (
	"fmt"
)

func main() {
	s1 := []int{5, 9, 12, 5, 24, 7, 8, 4, 17, 9, 2, 4, 13, 1, 24, 6, 2, 31, 12, 5, 10, 24}
	m1 := make(map[int]int)
	var s2 []int
	var max int
	var s3 []int

	for _, v := range s1 {
		if m1[v] != 0 {
			m1[v]++
		} else {
			m1[v] = 1
		}
	}
	fmt.Println(m1)

	for _, v := range m1 {
		s2 = append(s2, v)
	}
	max = s2[0]
	for i := 0; i < len(s2); i++ {
		if max < s2[i] {
			max = s2[i]
		}
	}
	fmt.Println(max)

	for k, v := range m1 {
		if v == max {
			s3 = append(s3, k)
		}
	}
	fmt.Println(s3)

}
