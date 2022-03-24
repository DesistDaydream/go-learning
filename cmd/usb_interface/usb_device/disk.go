package usbdevice

import "fmt"

// KingstonDisk 金士顿牌移动硬盘
type KingstonDisk struct {
	Name string
	Type string
	Data string
}

// NewKingstonDisk 实例化 KingstonDisk
func NewKingstonDisk() *KingstonDisk {
	return &KingstonDisk{
		Name: "A1",
		Type: "SSD",
		Data: "KingstonDisk fastest SSD",
	}
}

// Start 让 KingstonDisk 实现 USB 接口的方法
func (k *KingstonDisk) Start() {
	fmt.Println("金士顿SSD硬盘已连接")
}

// End 让 KingstonDisk 实现 USB 接口的方法
func (k *KingstonDisk) End() {
	//
}
