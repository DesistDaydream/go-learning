package main

import "fmt"

//1.定义一个接口
type USB interface {
	start()
	end()
}

//2.实现类
type Mouse struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标，准备就绪，可以开始工作，可以开始点点点。。。。。。")
}
func (m Mouse) end() {
	fmt.Println(m.name, "结束工作，可以安全退出。。")
}

type FlashDisk struct {
	name string
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "准备开始工作，可以进行数据存储了。。")
}
func (f FlashDisk) end() {
	fmt.Println(f.name, "可以弹出。。。")
}
func testInterface(usb USB) {
	usb.start()
	usb.end()
}
func main() {
	/*
	   接口：interface
	       方法的描述的集合。
	   实现类：只要实现了该接口中的所有的方法，那么该类就叫做该接口的实现类
	   注意点：
	       1.当需要接口类型的对象时，那么可以使用任意实现类对象代替。
	       2.接口对象不能访问实现类的属性。
	*/
	m := Mouse{"罗技小红"}
	fmt.Println(m.name)
	f := FlashDisk{"闪迪64G"}
	var usb USB //定义接口类型的变量
	usb = m     //创建该接口的任意实现类对象昂
	//fmt.Println(usb.name)//接口对象，不能访问实现类的属性
	//usb =f
	usb.start()
	usb.end()
	fmt.Println("--------------")
	testInterface(m)
	testInterface(f)
}
