package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//Nginx nginx  配置
type Nginx struct {
	Port    int    `yaml:"Port"`
	LogPath string `yaml:"LogPath"`
	Path    string `yaml:"Path"`
}

//Config   系统配置配置
type Config struct {
	Name      string `yaml:"SiteName"`
	Addr      string `yaml:"SiteAddr"`
	HTTPS     bool   `yaml:"Https"`
	SiteNginx Nginx  `yaml:"Nginx"`
}

func main() {
	var setting Config
	config, errRead := ioutil.ReadFile("./json_yaml/file/info.yaml")
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
}
