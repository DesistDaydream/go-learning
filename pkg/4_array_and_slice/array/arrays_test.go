package array

import (
	"testing"
)

func TestArrays(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "1.测试for循环对数组的赋值以及引用输出"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Arrays()
		})
	}
}

func TestArrInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "2.测试数组的初始化"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ArrInit()
		})
	}
}

func TestMultidimArray(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "3.测试多维数组"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MultidimArray()
		})
	}
}
