package main

import (
	"os"
	"text/template"
)

// FileParse 把文件当作模板主体进行解析的示例
func FileParse() {
	// ParseFiles 会执行 New 以及 Parse 方法，以文件名创建一个模板，并将文件内容定义到该模板中。
	t, err := template.ParseFiles("template/hello_world.tmpl")

	// 通过 Must 的验证，如果指定的文件无法找到，可以输出该信息。如果不验证，则无法看到无法找到文件的错误提示。
	t = template.Must(t, err)

	// 将指定的数据(第二个参数)应用于已解析的模板 t ，
	t.Execute(os.Stdout, "hello world")
}
