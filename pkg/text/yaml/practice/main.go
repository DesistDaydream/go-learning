package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config   系统配置配置
type Config struct {
	Name      string `yaml:"SiteName"`
	Addr      string `yaml:"SiteAddr"`
	HTTPS     bool   `yaml:"Https"`
	SiteNginx Nginx  `yaml:"Nginx"`
}

// Nginx nginx  配置
type Nginx struct {
	Port    int    `yaml:"Port"`
	LogPath string `yaml:"LogPath"`
	Path    string `yaml:"Path"`
}

func main() {
	filePath := "./test_files/text/info.yaml"
	var setting Config
	config, errRead := os.ReadFile(filePath)
	if errRead != nil {
		fmt.Print(errRead)
	}
	if err := yaml.Unmarshal(config, &setting); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(setting)
	fmt.Println(setting.Name)
	fmt.Println(setting.Addr)
	fmt.Println(setting.HTTPS)
	fmt.Println(setting.SiteNginx.Port)
	fmt.Println(setting.SiteNginx.LogPath)
	fmt.Println(setting.SiteNginx.Path)

	// TODO: 如何让程序生成的文件中，字段的值可以带双引号？
	setting.Name = "lichenhao"
	yamlData, _ := yaml.Marshal(&setting)
	os.WriteFile(filePath, yamlData, 0666)
}
