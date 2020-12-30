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
	var usbs = map[usbinterface.USB]bool{
		&usbdevice.KingstonDisk{}: true,
	}
	fmt.Println(usbs)
}
