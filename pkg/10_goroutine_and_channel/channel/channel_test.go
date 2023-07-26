package main

import (
	"testing"
)

// 通道的正确示例
func Test_channelCorrect(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "通道的正确示例"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channelCorrect()
		})
	}
}

// 通道会导致死锁的错误示例
// TODO: 有啥办法在测试时不卡主，而是直接 panic 吗？
// func Test_channelError(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{
// 		{name: "通道会导致死锁的错误示例"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			channelError()
// 		})
// 	}
// }

// 带缓冲的通道
func Test_channelBuffer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "带缓冲的通道"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			channelBuffer()
		})
	}
}

// 使用 range 关键字持续监听 Channel，并消费其对端生产的数据
func Test_forChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "使用 range 关键字持续监听 Channel，并消费其对端生产的数据"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rangeChannel()
		})
	}
}

func Test_isOkChannel(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "检测通道是否还有数据的语法"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isOkChannel()
		})
	}
}
