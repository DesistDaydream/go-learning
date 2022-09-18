package main

import "testing"

func TestGetStructAllInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "获取结构体中字段的 名称、Tag 等信息"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetStructAllInfo()
		})
	}
}
