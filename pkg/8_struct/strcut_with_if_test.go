package main

import (
	"fmt"
	"testing"
)

func TestNewStructWithInterface(t *testing.T) {
	got := NewStructWithFunction()
	fmt.Println(got)

}
