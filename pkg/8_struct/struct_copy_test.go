package main

import "testing"

func TestStructCopy(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "结构体之间的互相拷贝"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructCopy()
		})
	}
}
