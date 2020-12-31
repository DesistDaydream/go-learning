package main

import (
	"fmt"

	usbdevice "github.com/DesistDaydream/GoLearning/practice/usb_interface/usb_device"
	usbinterface "github.com/DesistDaydream/GoLearning/practice/usb_interface/usb_interface"
)

func main() {
	k := usbdevice.NewKingstonDisk()
	usbinterface.OperatorRead(k)
	m := usbdevice.NewMouse()
	usbinterface.OperatorRead(m)

	// usbs 测试接口多态效果，实现了接口的接口体，可以作为该接口的值。
	usbs := map[usbinterface.USB]bool{
		&usbdevice.KingstonDisk{}: true,
		&usbdevice.Mouse{}:        false,
	}
	fmt.Println(usbs)

	// 这是一个最简单的，把接口当作变量，把结构体当作值，然后调用接口下方法的例子
	// 此时 结构体=方法，所以作用在 usbVar 变量上的方法，实际上就是 func (k *KingstonDisk) Start() {}
	var usbsVar usbinterface.USB
	usbsVar = &usbdevice.KingstonDisk{}
	usbsVar.Start()
}
