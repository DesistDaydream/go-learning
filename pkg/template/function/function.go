package main

import (
	"os"
	"text/template"
)

func main() {
	tmplFile := "pkg/template/template_file/function.tmpl"
	// 模板中的函数示例
	t, err := template.ParseFiles(tmplFile)
	t = template.Must(t, err)

	// 将指定的数据(第二个参数)应用于已解析的模板 t ，
	t.Execute(os.Stdout, "hello world")
}
