package main

import (
	"testing"
)

func TestIncorrectUsageOfForrangeAndPointer(t *testing.T) {
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

func TestCorrectUsageOfForrangeAndPointer(t *testing.T) {
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

func TestCorrectUsageOfForAndPointer(t *testing.T) {
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
