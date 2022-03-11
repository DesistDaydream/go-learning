package main

import (
	"fmt"

	"github.com/DesistDaydream/go-learning/4_arrays_and_slices/array"
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

// 统计
func ConditionCount(combination []string, condition []string) bool {
	return array.IsSubset(condition, combination)
}

// isElement 正常统计
func isElement(combinations []string, condition string) bool {
	for _, combination := range combinations {
		if combination == condition {
			return true
		}
	}
	return false
}

// 判断遍历所有组合数的结果是否正确
func checkResult(n, k int, combinations [][]string) {
	rightCount := combination.Combination(n, k).Int64()
	if int(rightCount) == len(combinations) {
		fmt.Println("结果正确")
	} else {
		fmt.Println("结果错误，正确结果是：", rightCount)
	}

}

func main() {
	// 样本
	deck := []string{"a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "c", "c", "c", "d", "d", "e", "f", "g", "h", "i", "g", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	hand := []string{"a", "b", "a"}
	// deck := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "啊", "哦", "鱼"}
	// deck := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b", "b"}
	// deck := []string{"a", "b", "c", "d", "e", "f"}
	// deck := []string{"a", "a", "b", "b", "c", "e"}
	// 总体思路：列出 deck 样本中所有组合，对比每一个组合是否是 hand 的超集(即筛选出 hand 是 某些组合子集的组合)。将筛选出来的组合除以 deck 的总组合数，得出满足 hand 条件的组合比例

	var n int = len(deck)         // 样本中元素总数
	var k int = 5                 // 从样本中取出的元素数
	var TargetCombination int = 0 // 满足条件的组合数

	// 遍历样本，获取组合种类的列表
	combinations := ListCombinationKind(deck, combinationIndexs(n, int(k)))

	fmt.Println("原始组合总数:", len(combinations))
	// fmt.Println("原始组合列表:", combinations)
	checkResult(n, k, combinations)

	// 获取 deck 中，至少有 1 个 a 且 1 个 b 的组合数
	for _, combination := range combinations {
		// 下面的代码可以简化成递归处理
		for _, element1 := range combination {
			if element1 == "a" {
				for _, element2 := range combination {
					if element2 == "b" {
						TargetCombination++
						break
					}
				}
				// 若不退出，当数组中有多个 a 的时候，会匹配多次
				break
			}
		}

		if ConditionCount(combination, hand) {
			TargetCombination++
		}

		// 这种逻辑如何优化？
		// if isElement(combination, hand[0]) && isElement(combination, hand[1]) {
		// 	TargetCombination++
		// }
	}

	fmt.Println("指定组合总数:", TargetCombination)
	fmt.Println("取到指定组合的概率:", float64(TargetCombination)/float64(len(combinations)))
}
