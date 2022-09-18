package main

import (
	"fmt"
	"reflect"
)

type BaseData struct {
	// mysql
	DbUrl      string `yaml:"db_url" name:"数据库地址"`
	DbUser     string `yaml:"db_user" name:"数据库用户名"`
	DbPassWord string `yaml:"db_pass_word" name:"数据库密码"`
	DbName     string `yaml:"db_name" name:"数据库名"`
}

func GetStructAllInfo() {
	d := BaseData{
		DbUrl:      "url",
		DbUser:     "user",
		DbPassWord: "pw",
		DbName:     "name",
	}
	// 通过反射实例化一个类型
	t := reflect.TypeOf(d)
	// 通过反射实例化一个值
	v := reflect.ValueOf(d)
	for k := 0; k < t.NumField(); k++ {
		fmt.Println("name:", fmt.Sprintf("%+v", t.Field(k).Name), ", value:", fmt.Sprintf("%v", v.Field(k).Interface()), ", yaml:", t.Field(k).Tag.Get("yaml"))
	}
}
