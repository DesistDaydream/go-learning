package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// 将字段名改为大写开头
type person struct {
	Name string
	Age  int
}

type house struct {
	Country string
	City    string
	State   string
}

func multiStructUnifiedProcessing() {
	var (
		my      string = "DesistDaydream|18"
		myHouse string = "China|Tianjin|Tianjin"
	)

	// 创建结构体实例
	p := &person{}
	h := &house{}

	// 处理数据并填充到结构体中
	if err := processing(my, p); err != nil {
		fmt.Printf("Error processing person: %v\n", err)
	}
	if err := processing(myHouse, h); err != nil {
		fmt.Printf("Error processing house: %v\n", err)
	}

	// 打印结果
	fmt.Printf("Person: %+v\n", p)
	fmt.Printf("House: %+v\n", h)
}

func processing(text string, dst interface{}) error {
	// 分割字符串
	parts := strings.Split(text, "|")

	// 获取目标结构体的反射值和类型
	v := reflect.ValueOf(dst).Elem()
	t := v.Type()

	// 检查字段数量是否匹配
	if t.NumField() != len(parts) {
		return fmt.Errorf("field count mismatch: struct has %d fields but got %d values", t.NumField(), len(parts))
	}

	// 遍历结构体的字段
	for i := 0; i < t.NumField(); i++ {
		field := v.Field(i)
		// 根据字段类型进行相应的转换和赋值
		switch field.Kind() {
		case reflect.String:
			field.SetString(parts[i])
		case reflect.Int:
			val, err := strconv.Atoi(parts[i])
			if err != nil {
				return fmt.Errorf("failed to convert %s to int: %v", parts[i], err)
			}
			field.SetInt(int64(val))
		// 可以根据需要添加其他类型的处理
		default:
			return fmt.Errorf("unsupported field type: %v", field.Kind())
		}
	}

	return nil
}

// func main() {
// 	multiStructProcessing()
// }
