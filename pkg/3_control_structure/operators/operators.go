package main

import (
	"fmt"
	"strconv"
)

// 定义网络接口标志常量
const (
	IFF_UP          = 1 << 0
	IFF_BROADCAST   = 1 << 1
	IFF_DEBUG       = 1 << 2
	IFF_LOOPBACK    = 1 << 3
	IFF_POINTOPOINT = 1 << 4
	IFF_NOTRAILERS  = 1 << 5
	IFF_RUNNING     = 1 << 6
	IFF_NOARP       = 1 << 7
	IFF_PROMISC     = 1 << 8
	IFF_ALLMULTI    = 1 << 9
	IFF_MASTER      = 1 << 10
	IFF_SLAVE       = 1 << 11
	IFF_MULTICAST   = 1 << 12
	IFF_PORTSEL     = 1 << 13
	IFF_AUTOMEDIA   = 1 << 14
	IFF_DYNAMIC     = 1 << 15
)

// 映射标志到其名称
var flagNames = map[int]string{
	IFF_UP:          "IFF_UP",
	IFF_BROADCAST:   "IFF_BROADCAST",
	IFF_DEBUG:       "IFF_DEBUG",
	IFF_LOOPBACK:    "IFF_LOOPBACK",
	IFF_POINTOPOINT: "IFF_POINTOPOINT",
	IFF_NOTRAILERS:  "IFF_NOTRAILERS",
	IFF_RUNNING:     "IFF_RUNNING",
	IFF_NOARP:       "IFF_NOARP",
	IFF_PROMISC:     "IFF_PROMISC",
	IFF_ALLMULTI:    "IFF_ALLMULTI",
	IFF_MASTER:      "IFF_MASTER",
	IFF_SLAVE:       "IFF_SLAVE",
	IFF_MULTICAST:   "IFF_MULTICAST",
	IFF_PORTSEL:     "IFF_PORTSEL",
	IFF_AUTOMEDIA:   "IFF_AUTOMEDIA",
	IFF_DYNAMIC:     "IFF_DYNAMIC",
}

// 检查特定标志是否被设置
func checkFlag(flags int, flag int) bool {
	return flags&flag == flag
}

// 解析标志
func parseFlags(hexFlags string) []string {
	flags, _ := strconv.ParseInt(hexFlags, 0, 32)
	var activeFlags []string

	for flag, name := range flagNames {
		if checkFlag(int(flags), flag) {
			activeFlags = append(activeFlags, name)
		}
	}

	return activeFlags
}

// 模拟 Linux 如何判断一个网络设备的 Flags 都设置了哪些，
func unixNetDevFlagsDemo() {
	// 从 /sys/class/net/<interface>/flags 读取标志值
	hexFlags := "0x1003"
	fmt.Printf("Flags value: %s\n", hexFlags)

	activeFlags := parseFlags(hexFlags)
	fmt.Println("Active flags:")
	for _, flag := range activeFlags {
		fmt.Printf("- %s\n", flag)
	}
}

func sampleDemo() {
	// 定义一些标志位
	const (
		FLAG_A = 1 << 0 // 0001
		FLAG_B = 1 << 1 // 0010
		FLAG_C = 1 << 2 // 0100
	)

	flags := FLAG_A | FLAG_B
	fmt.Printf("Flags: %b\n", flags)

	fmt.Printf("Flags: %b\n", 3)
	fmt.Printf("Flags: %b\n", 3<<2)
	fmt.Printf("Flags: %d\n", 3<<2)
}
