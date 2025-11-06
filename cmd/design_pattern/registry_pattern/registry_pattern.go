package main

import "fmt"

type demoStruct struct {
	name string
	age  int
}

func exportDataToPDF(demoData demoStruct, format string) {
	fmt.Printf("导出数据到 PDF 格式: %+v\n", demoData)
}

func exportDataToCSV(demoData demoStruct, format string) {
	fmt.Printf("导出数据到 CSV 格式: %+v\n", demoData)
}

func exportDataToJSON(demoData demoStruct, format string) {
	fmt.Printf("导出数据到 JSON 格式: %+v\n", demoData)
}

func exportData(demoData demoStruct, format string) {
	// 方式一：传统的 if else
	// 如果使用传统的 if else，当有新增的导出格式时，需要修改 exportData 函数的逻辑。
	if format == "pdf" {
		exportDataToPDF(demoData, format)
	} else if format == "csv" {
		exportDataToCSV(demoData, format)
	} else if format == "json" {
		exportDataToJSON(demoData, format)
	}

	// 方式二：设计模式，注册模式
	// 使用**注册模式**后，如果有新增的导出格式，只需要新增一个导出函数，然后注册到 map 中即可。
	// 而不需要修改 exportData 函数的逻辑。
	exportFunc, ok := exportDataRegistry[format]
	if !ok {
		fmt.Printf("不支持 %s 格式的导出\n", format)
		return
	}
	exportFunc(demoData, format)
}

// ######## 注册模式 ########
// 1. 定义一个全局的 map，key 为格式，value 为导出函数
// 2. 导出函数注册到 map 中
// 3. 导出函数从 map 中获取，根据格式调用对应的导出函数
// ######## 注册模式 ########
var exportDataRegistry = map[string]func(demoStruct, string){
	"pdf":  exportDataToPDF,
	"csv":  exportDataToCSV,
	"json": exportDataToJSON,
}

// https://www.bilibili.com/video/BV113WszsEoc
func main() {
	demoData := demoStruct{
		name: "张三",
		age:  18,
	}

	exportData(demoData, "pdf")
	exportData(demoData, "csv")
	exportData(demoData, "json")
}
