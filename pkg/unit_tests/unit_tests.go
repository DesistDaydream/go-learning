package main

// 一个带有参数和返回值的测试用例示例
func UnitTests(needArgs string) bool {
	if needArgs == "unittests" {
		return true
	} else {
		return false
	}
}
