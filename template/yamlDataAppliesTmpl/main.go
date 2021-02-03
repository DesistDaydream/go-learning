package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

func main() {
	// 创建并解析模板文件
	t := template.Must(template.ParseFiles("./template/yaml.tmpl"))

	// ReadFile 读取 yaml 文件内容，并将内容放入 config 中
	config, errRead := ioutil.ReadFile("./template/info.yaml")
	if errRead != nil {
		fmt.Print(errRead)
	}

	// make 创建一个存储 yaml 数据的地方，通过这种方式，不用关联结构体，直接将 yaml 数据传到模板中。
	yamlInfo := make(map[interface{}]interface{})

	// Unmarshal 解析 yaml 格式数据，并将解析结果放入 yamlInfo 中。
	errUnmarshal := yaml.Unmarshal(config, &yamlInfo)
	if errUnmarshal != nil {
		log.Fatalf("error: %v", errUnmarshal)
	}

	// Execute 将处理好的所有 yaml 文件中的内容，应用于模板，并输出。
	t.Execute(os.Stdout, yamlInfo)
}
