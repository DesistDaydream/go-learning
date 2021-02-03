package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: 
  - 3
  - 4
`

// B 结构体，字段注释写不写都行
type B struct {
	RenamedC int   `yaml:"c"`
	D        []int `yaml:",flow"`
}

// T 结构体
// Note: struct 字段必须是全局生效的，以便可以正确 unmarshal 该结构体并讲 yaml 数据填充进来
type T struct {
	A string
	B
}

func main() {
	// 声明结构体，该结构体用于存储 yaml 数据。
	// 这种方式存储的 yaml 数据，单行 yaml 格式无法显示 key
	t := T{}

	// Unmarshal 具有打散、解包等含义。将参数2中的数据进行解码后，传给参数1。
	// 官方文档称 Unmarshel 的行为 decode(解码) yaml 数据。
	// 用白话说就是将多行 yaml 格式合并为1行 yaml 格式。
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- 单行 yaml 格式:\n%v\n", t)
	fmt.Printf("注意：使用 Unmarshal 处理后的的数据常用于进行值传递，比如将这种解析后的数据传递到 go 模板、或者传递到 kubeclt 后进行二次加工并输出等等\n\n")

	// Marshel 具有整理、排列、打包、编组等含义。将参数中的数据进行编码，并返回序列化后的 yaml 数据。
	// 官方文档称 Marshel 的行为 encode(编码) yaml 数据。
	// Marshel() 参数可以接受的数据格式为 `(...) yaml:"[<key>][,<flag1>[,<flag2>]]" (...)`
	d1, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- 多行 yaml 格式:\n%s\n", string(d1))

	// 通过 make 创建一个存储 yaml 数据的地方
	// 这种方式存储的 yaml 数据，单行 yaml 格式可以显示 key
	m := make(map[interface{}]interface{})

	// Unmarshal 的作用与上面一样
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- 单行 yaml 格式:\n%v\n\n", m)

	// Marshal 的作用与上面一样
	d2, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- 多行 yaml 格式:\n%s\n", string(d2))
}
