package array

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("1.测试for循环对数组的赋值以及引用输出")
	arrays()

	fmt.Println("\n2.测试数组的初始化")
	arrInit()

	fmt.Println("\n3.测试多维数组")
	multidim_array()

	superset := []string{"z", "d", "a", "c", "b", "c"}
	subset1 := []string{"a", "c", "b"}
	subset2 := []string{"a", "c", "c"}
	result1 := IsSubset(subset1, superset)
	result2 := IsSubset(subset2, superset)
	fmt.Printf("命题: %v 是 %v 的子集.\n结论: %v\n", subset1, superset, result1)
	fmt.Printf("命题: %v 是 %v 的子集.\n结论: %v\n", subset2, superset, result2)
}
