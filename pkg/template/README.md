# Go Template 标准库
Go 原生支持的模板分为两种
1. [text](https://pkg.go.dev/text/template)
2. [html](https://pkg.go.dev/html/template)

参考文章：  
[Go template 用法详解](https://www.cnblogs.com/f-ck-need-u/p/10053124.html)  
[深入剖析 Go template](https://www.cnblogs.com/f-ck-need-u/p/10035768.html)  

# text/template 库
template库 实现了数据驱动的用于生成文本输出的模板。

如果要生成HTML格式的输出，参见html/template包，该包提供了和本包相同的接口，但会自动将输出转化为安全的HTML格式输出，可以抵抗一些网络攻击。

通过将模板应用于一个数据结构（即该数据结构作为模板的参数）来执行，来获得输出。模板中的注释引用数据接口的元素（一般如结构体的字段或者字典的键）来控制执行过程和获取需要呈现的值。模板执行时会遍历结构并将指针表示为 `.` (称之为"dot")指向运行过程中数据结构的当前位置的值。

用作模板的输入文本必须是utf-8编码的文本。"Action" 数据运算和控制单位由 `{{` 和 `}}`界定；在Action之外的所有文本都不做修改的拷贝到输出中。Action内部不能有换行，但注释可以有换行。

经解析生成模板后，一个模板可以安全的并发执行。

下面是一个简单的例子，可以打印"17 of wool"。
```
type Inventory struct {
	Material string
	Count    uint
}
sweaters := Inventory{"wool", 17}
tmpl, err := template.New("test").Parse("{{.Count}} of {{.Material}}")
if err != nil { panic(err) }
err = tmpl.Execute(os.Stdout, sweaters)
if err != nil { panic(err) }
```
更复杂的例子在下面。

## Actions

### Arguments

### Pipelines