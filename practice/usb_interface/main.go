package main

import (
	"fmt"

	usbinterface "github.com/DesistDaydream/GoLearning/practice/usb_interface/usb_interface"
)

// KingstonDisk 金士顿牌移动硬盘
type KingstonDisk struct {
	Name string
	Type string
	Data string
}

// NewKingstonDisk is
func NewKingstonDisk() *KingstonDisk {
	return &KingstonDisk{
		Name: "A1",
		Type: "SSD",
		Data: "KingstonDisk fastest SSD",
	}
}

// Start is
func (k *KingstonDisk) Start() {
	fmt.Println("金士顿SSD硬盘已连接")
}

// End is
func (k *KingstonDisk) End() {

}

func main() {
	//
	k := NewKingstonDisk()
	usbinterface.OperatorRead(k)
}
