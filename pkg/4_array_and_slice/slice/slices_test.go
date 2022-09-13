package main

import (
	"testing"
)

func TestGoSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "切片的基本用法"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GoSlice()
		})
	}
}

func TestSliceReslice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "切片的重组"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceReslice()
		})
	}
}

func TestSliceAppend(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "切片的追加(append)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceAppend()
		})
	}
}

func TestSliceCopy(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "切片的复制(copy)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceCopy()
		})
	}
}

func Test_sliceString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "字符串、数组和切片的应用"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceString()
		})
	}
}
