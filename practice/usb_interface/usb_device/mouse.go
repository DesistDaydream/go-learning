package usbdevice

import "fmt"

// Mouse 一个鼠标
type Mouse struct {
	Name string
	Data string
}

// NewMouse 实例化 Mouse
func NewMouse() *Mouse {
	return &Mouse{
		Name: "Razer",
		Data: "This is Razer Mouse",
	}
}

// Start 让 Mouse 实现 USB 接口的方法
func (k *Mouse) Start() {
	fmt.Println("鼠标已连接")
}

// End 让 Mouse 实现 USB 接口的方法
func (k *Mouse) End() {
	//
}
