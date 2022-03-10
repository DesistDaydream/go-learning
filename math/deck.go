package main

import (
	"fmt"
	"time"

	"github.com/DesistDaydream/go-learning/math/combination"
)

//组合算法(从nums中取出m个数)
func combinationIndexs(n int, k int) [][]int {
	if k < 1 || k > n {
		fmt.Println("Illegal argument. Param m must between 1 and len(nums).")
		return [][]int{}
	}

	//保存最终结果的数组，总数直接通过数学公式计算
	result := make([][]int, 0, combination.Combination(n, k).Int64())
	//保存每一个组合的索引的数组，1表示选中，0表示未选中
	indexs := make([]int, n)
	for i := 0; i < n; i++ {
		if i < k {
			indexs[i] = 1
		} else {
			indexs[i] = 0
		}
	}

	//第一个结果
	result = addTo(result, indexs)
	for {
		find := false
		//每次循环将第一次出现的 1 0 改为 0 1，同时将左侧的1移动到最左侧
		for i := 0; i < n-1; i++ {
			if indexs[i] == 1 && indexs[i+1] == 0 {
				find = true

				indexs[i], indexs[i+1] = 0, 1
				if i > 1 {
					moveOneToLeft(indexs[:i])
				}
				result = addTo(result, indexs)

				break
			}
		}

		//本次循环没有找到 1 0 ，说明已经取到了最后一种情况
		if !find {
			break
		}
	}

	return result
}

//将ele复制后添加到arr中，返回新的数组
func addTo(arr [][]int, ele []int) [][]int {
	newEle := make([]int, len(ele))
	copy(newEle, ele)
	arr = append(arr, newEle)

	return arr
}

func moveOneToLeft(leftNums []int) {
	//计算有几个1
	sum := 0
	for i := 0; i < len(leftNums); i++ {
		if leftNums[i] == 1 {
			sum++
		}
	}

	//将前sum个改为1，之后的改为0
	for i := 0; i < len(leftNums); i++ {
		if i < sum {
			leftNums[i] = 1
		} else {
			leftNums[i] = 0
		}
	}
}

//根据索引号数组得到元素数组
func ListCombinationKind(nums []string, indexs [][]int) [][]string {
	if len(indexs) == 0 {
		return [][]string{}
	}

	result := make([][]string, len(indexs))

	for i, v := range indexs {
		line := make([]string, 0)
		for j, v2 := range v {
			if v2 == 1 {
				line = append(line, nums[j])
			}
		}
		result[i] = line
	}

	return result
}

var (
	// 满足条件的组合数
	targetCombination int = 0
)

/*
【排列组合问题：n个数中取k个】
*/
func main() {
	// 样本
	deck := []string{"a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "c", "c", "c", "d", "d", "e", "f", "g", "h", "i", "g", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	// deck := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "啊", "哦", "鱼"}
	// deck := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}
	// deck := []string{"a", "b", "c", "d", "e", "f"}
	// deck := []string{"a", "a", "b", "b", "e", "f"}

	// 数组长度
	var n int = len(deck)
	// 取出元素个数
	var k int = 5

	timeStart := time.Now()
	indexs := combinationIndexs(n, int(k))
	// 组合种类的列表
	combinations := ListCombinationKind(deck, indexs)
	timeEnd := time.Now()

	fmt.Println("原始组合总数:", len(combinations))
	// fmt.Println("所有组合结果:", combinationKinds)
	fmt.Println("时间消耗:", timeEnd.Sub(timeStart))

	// 判断结果是否正确
	rightCount := combination.Combination(n, k).Int64()
	if int(rightCount) == len(combinations) {
		fmt.Println("结果正确")
	} else {
		fmt.Println("结果错误，正确结果是：", rightCount)
	}

	// 获取 deck 中，至少有 1 个 a 且 1 个 b 的组合数
	// for i, combination := range combinations {
	// 	for i1, condition1 := range combination {
	// 		if condition1 == "a" {
	// 			for i2, condition2 := range combination {
	// 				if condition2 == "b" {
	// 					targetCombination++
	// 					fmt.Printf("原始索引: %v, 条件1索引: %v, 条件2索引: %v\n", i, i1, i2)
	// 					_, _, _ = i, i1, i2
	// 					break
	// 				}
	// 			}
	// 			// 若不退出，当数组中有多个 a 的时候，会匹配多次
	// 			break
	// 		}
	// 	}
	// }
	hand := []string{"a", "b", "c", "d", "e"}
	for _, combinationKind := range combinations {
		ConditionCount(combinationKind, hand)
	}

	fmt.Println("指定组合总数:", targetCombination)
	fmt.Println("取到指定组合的概率:", float64(targetCombination)/float64(len(combinations)))
}

func ConditionCount(combinations []string, condition []string) bool {
	for _, combination := range combinations {
		if len(condition) == 0 {
			return true
		}

		if condition[0] == combination {
			if len(condition) == 1 {
				targetCombination++
			}
			return ConditionCount(combinations, condition[1:])
		}
	}
	return false
}
