package main

import (
	"os"
	"strings"
	"text/template"
)

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	// 创建可以用在 Go tempalte 中的自定义函数
	var tmplFunc = template.FuncMap{
		"toUpper": toUpper,
	}

	// Notes: Funcs() 必须在各种 ParseXXX() 用于解析模板主体的函数之前调用
	// 所一要先 New()，然后 Funcs()，最后 ParseFiles()
	// New() 的参数是模板的名字，要与 ParseFiles() 参数中的文件名一致
	t := template.Must(template.New("function_custom.tmpl").Funcs(tmplFunc).ParseFiles("pkg/template/template_file/function_custom.tmpl"))

	t.Execute(os.Stdout, "hello world")
}
