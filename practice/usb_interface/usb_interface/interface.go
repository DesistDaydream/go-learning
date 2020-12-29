package usbinterface

import "fmt"

// USB 接口
type USB interface {
	Start()
	End()
}

// OperatorRead 启动插入接口的设备并从中读取信息读取、读取后结束
func OperatorRead(u USB) {
	u.Start()
	fmt.Printf("当前连接的设备信息为：%v\n", u)
	u.End()
}
