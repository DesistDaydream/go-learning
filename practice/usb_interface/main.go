package main

import (
	usbdevice "github.com/DesistDaydream/GoLearning/practice/usb_interface/usb_device"
	usbinterface "github.com/DesistDaydream/GoLearning/practice/usb_interface/usb_interface"
)

func main() {
	//
	k := usbdevice.NewKingstonDisk()
	usbinterface.OperatorRead(k)
}
