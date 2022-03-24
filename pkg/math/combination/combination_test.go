package combination

import (
	"log"
	"testing"
)

func TestCombination(t *testing.T) {
	var n int = 40
	var k int = 5
	var a int = 6
	var b int = 7

	log.Printf("从 %v 个元素中取 %v 个元素的组合数:%v", n, k, Combination(n, k))

	fenzi := Combination(n-b, k).Int64() + Combination(n-a, k).Int64() - Combination(n-a-b, k).Int64()
	fenmu := Combination(n, k).Int64()

	log.Printf("分子：%v 分母：%v", float64(fenzi), float64(fenmu))

	result := float64(fenzi) / float64(fenmu)
	log.Printf("取到指定组合的概率:%v", 1-result)
}
