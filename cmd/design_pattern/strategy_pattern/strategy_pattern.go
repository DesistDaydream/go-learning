package main

import "fmt"

// 策略模式
func main() {
	a := 1
	b := 2
	c := 3
	var (
		conditionOne   bool = func() bool { return a == 1 }()
		conditionTwo   bool = func() bool { return b == 2 }()
		conditionThree bool = func() bool { return c == 3 }()
	)

	if conditionOne {
		fmt.Println("满足条件 1 后执行")
		if conditionTwo {
			fmt.Println("满足条件 2 后执行")
		}
	} else if conditionThree {
		fmt.Println("满足条件 3 后执行")
	} else {
		fmt.Println("都不满足后执行")
	}

	// 使用设计模式实现上述 if else 逻辑
	strategyPattern()
}

// 策略模式就是将 if else 中每个条件判断抽象为一个 对象（i.e. go 的 struct）。
// 这些 条件对象 通常要包含两个方法: 1. 判断策略是否应该执行 2. 策略执行的具体逻辑
//
// 多个 条件对象 要实现一个接口。
type Strategy interface {
	IsMatch(data *StrategyData) bool // 1. 判断策略是否应该执行。相当于 if else 中的条件判断
	Matched(data *StrategyData)      // 2. 策略执行的具体逻辑。相当于 if else 中的具体逻辑
	// ... 其他方法，比如：优先级排序、条件判断的依赖关系、策略失败时的处理、etc.
	NotMatched(data *StrategyData) // 3. 策略不执行时的处理
}

// 通常一个方法里应该包含一个策略数据的参数
type StrategyData struct {
	a int
	b int
	c int
}

// ######## 定义 条件对象 作为策略 ########
// ！！！如果需要添加新的策略，只需要创建实现 Strategy 接口的 struct ，并将 struct 添加到下面的 StrategyManager() 中即可（i.e. 实例化时添加新策略）！！！
//
// 策略1
type StrategyOne struct{}

func (c *StrategyOne) IsMatch(data *StrategyData) bool {
	return data.a == 1
}
func (c *StrategyOne) Matched(data *StrategyData) {
	fmt.Println("策略模式满足条件 1 后执行")
}
func (c *StrategyOne) NotMatched(data *StrategyData) {
	fmt.Println("策略模式不满足条件 1 后执行")
}

// 策略2
type StrategyTwo struct{}

func (c *StrategyTwo) IsMatch(data *StrategyData) bool {
	return data.b == 2
}
func (c *StrategyTwo) Matched(data *StrategyData) {
	fmt.Println("策略模式满足条件 2 后执行")
}
func (c *StrategyTwo) NotMatched(data *StrategyData) {
	fmt.Println("策略模式不满足条件 2 后执行")
}

// 策略3
type StrategyThree struct{}

func (c *StrategyThree) IsMatch(data *StrategyData) bool {
	return data.c == 3
}
func (c *StrategyThree) Matched(data *StrategyData) {
	fmt.Println("策略模式满足条件 3 后执行")
}
func (c *StrategyThree) NotMatched(data *StrategyData) {
	fmt.Println("策略模式不满足条件 3 后执行")
}

// ######## 定义策略结束 ########

// 定义一个策略管理器，管理多个策略
type StrategyManager struct {
	strategies []Strategy
}

// 实例化一个策略管理器，为其添加多个策略
func NewStrategyManager() *StrategyManager {
	return &StrategyManager{
		strategies: []Strategy{
			&StrategyOne{},
			&StrategyTwo{},
			&StrategyThree{},
		},
	}
}

// Process 处理策略的逻辑。
// 在这里遍历所有策略，找到匹配的策略并执行。就像 if-else 也要一个一个判断一样。
// 但是，如果策略很多（if else 条件很多），if else 会显得很臃肿，尤其对于嵌套的 if else，策略模式会让代码看起来更清晰。整体条件判断的可扩展性也很强。
// 这里说的清晰是指：没有 if else 那种频繁缩进，每个 if 下的具体逻辑都在独立的方法中。
func (m *StrategyManager) Process(data *StrategyData) {
	for _, strategy := range m.strategies {
		// 判断每个策略是否应该执行
		if strategy.IsMatch(data) {
			// 策略匹配，则执行该策略对应的逻辑
			strategy.Matched(data)
		} else {
			// 策略不匹配，则执行该策略对应的逻辑
			strategy.NotMatched(data)
		}
	}
}

// 使用策略模式
func strategyPattern() {
	processor := NewStrategyManager()
	processor.Process(&StrategyData{a: 1, b: 2, c: 4})
}
