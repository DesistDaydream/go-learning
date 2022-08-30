package main

import (
	"testing"
)

func TestForRange与指针的错误用法(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "for range与指针的错误用法",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IncorrectUsageOfForrangeAndPointer()
		})
	}
}

func TestForRange与指针的正确用法一(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "for range与指针的正确用法一"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CorrectUsageOfForrangeAndPointer()
		})
	}
}

func TestFor与指针的正确用法(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{name: "for range与指针的正确用法二"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CorrectUsageOfForAndPointer()
		})
	}
}
