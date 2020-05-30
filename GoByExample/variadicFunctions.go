package main

import (
	"fmt"
	"reflect"
)
// 这个函数接受任意数目的int类型的参数
func sum(nums ...int) {
	fmt.Print(nums, " ")
	fmt.Println(reflect.TypeOf(nums))
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	// 变参函数使用常规的调用方式，传入独立的参数
	// 变参函数可以传递任意数量的参数
	sum(1, 2)
	sum(1, 2, 3)
	// 如果有一个多个值的slice，想把它们作为参数使用，需要这样调用`FunctionID(SliceID...)`
	// 注意3个点不要忘记写
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
