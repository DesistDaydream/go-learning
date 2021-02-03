package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

// User 用户信息
type User struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
}

//Nginx nginx  配置
type Nginx struct {
	Port    int    `yaml:"port"`
	LogPath string `yaml:"logPath"`
	Path    string `yaml:"path"`
}

//YamlInfo yaml文件信息
type YamlInfo struct {
	Name  string `yaml:"name"`
	User  []User `yaml:"user"`
	Nginx Nginx  `yaml:"nginx"`
}

func main() {
	// 创建并解析模板文件
	t := template.Must(template.ParseFiles("../template/struct_store_yamlData.tmpl"))

	// 读取 yaml 文件内容，并将内容放入 config 中后
	config, errRead := ioutil.ReadFile("../template/info.yaml")
	if errRead != nil {
		fmt.Print(errRead)
	}

	// 声明结构体，用于存储 yaml 格式的数据。
	var yamlInfo YamlInfo

	// 通过 Unmarshal 解析 yaml 格式数据，并将解析结果放入 yamlInfo 中。
	errUnmarshal := yaml.Unmarshal(config, &yamlInfo)
	if errUnmarshal != nil {
		log.Fatalf("error: %v", errUnmarshal)
	}

	// 将处理好的所有 yaml 文件中的内容，应用于模板，并输出。
	t.Execute(os.Stdout, yamlInfo)
}
