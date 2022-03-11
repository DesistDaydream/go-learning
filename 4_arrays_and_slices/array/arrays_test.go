package array

import (
	"fmt"
	"testing"
)

func TestArrays(t *testing.T) {
	fmt.Println("1.测试for循环对数组的赋值以及引用输出")
	Arrays()
}

func TestArrInit(t *testing.T) {
	fmt.Println("2.测试数组的初始化")
	ArrInit()
}

func TestMultidimArray(t *testing.T) {
	fmt.Println("3.测试多维数组")
	MultidimArray()
}

func TestIsSubset(t *testing.T) {
	superset := []string{"z", "d", "a", "c", "b", "c"}
	subset1 := []string{"a", "c", "b"}
	subset2 := []string{"a", "c", "c"}
	result1 := IsSubset(subset1, superset)
	result2 := IsSubset(subset2, superset)
	fmt.Printf("命题: %v 是 %v 的子集.\n结论: %v\n", subset1, superset, result1)
	fmt.Printf("命题: %v 是 %v 的子集.\n结论: %v\n", subset2, superset, result2)
}
