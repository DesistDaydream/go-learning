package combination

import (
	"log"
	"math/big"
)

func InitDeck() {
	// deck := []string{"a", "a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b", "b", "c", "d", "e", "f", "g", "h", "i", "g", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "啊", "哦", "鱼"}
	// log.Printf("长度：%v", len(deck))
	// log.Printf("随机取出的一个数：%v", deck[index])
	// for i := 0; i < 5; i++ {
	// 	n, _ := rand.Int(rand.Reader, big.NewInt(39))
	// 	println(deck[n.Int64()])
	// }
}

// 阶乘
func Factorial(n uint64) uint64 {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}

// 大数阶乘
func BigFactorial(n *big.Int) *big.Int {
	if n.Int64() == 1 {
		return big.NewInt(1)
	} else {
		return n.Mul(n, BigFactorial(big.NewInt(n.Int64()-1)))
	}
}

// 组合
func Combination(n, k int) *big.Int {
	result := big.NewInt(0)
	if n <= k {
		log.Println(n, k)
		return result
	}
	nF := BigFactorial(big.NewInt(int64(n)))
	kF := BigFactorial(big.NewInt(int64(k)))
	nkF := BigFactorial(big.NewInt(int64(n - k)))
	return result.Div(result.Div(nF, kF), nkF)
}
