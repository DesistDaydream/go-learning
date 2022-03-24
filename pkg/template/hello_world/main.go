package main

import (
	"os"
	"text/template"
)

func main() {
	// New 创建一个名为 hello world 的未定义模板。此时模板中无任何内容，模板内容以及要传递的值都没有，仅仅是生成一个模板指针，用来后续调用。
	// 这种行为类似于 open 一个文件，或者 new 一个数组之类，或者 open 一个数据库，或者 open 一个 redis 等等。
	t := template.New("hello world")

	// Parse 定义模板 t 。将 Parse() 方法中的参数解析为模板的主体。
	// 如果用一个文件来类比的话，现在就是有一个名为"hello world"的文件，这个文件的内容为"这是一个 {{ . }} 模板"
	t, _ = t.Parse("这是一个最简单的 {{ . }} 模板示例\n")

	// 上述两个步骤可以合并为一行代码，示例如下
	// 这是一个没有 Must 验证的代码
	// t, _ := template.New("hello world").Parse("这是一个 {{ . }} 模板")

	// Must 这是一个帮助程序，用来对模板 t 进行验证，并不影响后续对模板的输出，可省略。这种验证对于简单的代码没有什么意义。
	// t = template.Must(t, err)

	// FileParse 本目录下的另一个文件中的函数，用来测试解析一个文件
	// Must 验证的基本功能可以参考 FilePparse() 中 Must 的用法。注意：同 package 下不同文件调用函数，需要使用 go run *.go 命令，让所有 go 文件都处于此次运行指令中。
	FileParse()

	// 将指定的数据(第二个参数)应用于已解析的模板 t ,并将输出写入 第一个参数 中。
	// 这是一个合并、替换的过程。将模板中的模板语句(这种语句称为 Pipeline，在这里是 {{ . }})，替换成 第二个参 数传入的数据。并与 Pipeline 之外的语句合并，最后写入到 第一个参数 中。
	t.Execute(os.Stdout, "hello world")
}
