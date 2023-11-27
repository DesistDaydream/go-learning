package main

import (
	"testing"
)

func Test_dataType(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "数据类型检查"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataType()
		})
	}
}

func Test_typeConversions(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "数据类型转换"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typeConversions()
		})
	}
}

func Test_stringIntFloat(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "使用 strconv 标准库来实现字符串类型与其他基本数据类型之间的类型转换。"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringIntFloat()
		})
	}
}

func Test_unicode(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "字符串转Unicode，Unicode转字符串"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unicode()
		})
	}
}
