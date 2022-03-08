package main

import (
	"log"
	"math/big"
)

// 阶乘
func Factorial(n uint64) uint64 {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}

// 大数阶乘
func BigFactorial(s *big.Int) *big.Int {
	if s.Int64() == 1 {
		return big.NewInt(1)
	} else {
		return s.Mul(s, BigFactorial(big.NewInt(s.Int64()-1)))
	}
}

// 组合
func Combination(n, k int64) *big.Int {
	result := big.NewInt(0)
	nF := BigFactorial(big.NewInt(n))
	kF := BigFactorial(big.NewInt(k))
	nkF := BigFactorial(big.NewInt(n - k))
	// log.Printf("%d 的阶乘是 %d\n", n, nF)
	// log.Printf("%d 的阶乘是 %d\n", k, kF)
	// log.Printf("%d 的阶乘是 %d\n", n-k, nkF)
	return result.Div(result.Div(nF, kF), nkF)
}

func main() {
	// // deck := []string{"a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "c", "d", "e", "f", "g", "h", "i", "g", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "啊", "哦", "鱼"}
	// deck := []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	// log.Printf("长度：%v", len(deck))
	// // log.Printf("随机取出的一个数：%v", deck[index])
	// for i := 0; i < 5; i++ {
	// 	n, _ := rand.Int(rand.Reader, big.NewInt(39))
	// 	println(deck[n.Int64()])
	// }

	var n int64 = 40
	var k int64 = 5

	log.Printf("从 %v 个元素中取 %v 个元素的组合数:%v", n, k, Combination(n, k))

	fenzi := Combination(33, 5).Int64() + Combination(34, 5).Int64() - Combination(27, 5).Int64()
	fenmu := Combination(40, 5).Int64()

	result := float64(fenzi) / float64(fenmu)
	log.Printf("%v", 1-result)

}
