# 实验

1. [字符串的处理](#字符串的处理)
1. [读取用户的输入](#读取用户的输入)
1. [文件读写](#文件读写)
1. [命令行参数处理](#命令行参数处理)
1. [启动外部命令和程序](#启动外部命令和程序)

# 读写数据

golang 对于读写的操作最基础的是使用 bufio 这个 package。bufio 有一些个方法可以实现读取文件、标准输入等等操作，还可以进行输出等。

## 读取用户的输入

大多数的程序都是处理输入，产生输出；这也正是`计算`的定义。但是程序如何获取要处理的输入数据呢？有一些程序生成自己的数据，但是通常情况下，输入来自于程序外部，e.g.文件、、网络连接、其他程序的输出、敲键盘的用户、命令行参数或其它类似的输入源。想要使用 Go 语言的输入输出功能，一般不外乎下面 3 步

1. 获取输入源的定位符，e.g.文件描述符、用户的标准输入等
2. 通过输入源的定位符，把输入内容放到缓冲区并充缓冲器发送给变量
3. 打印缓冲区的变量即可实现输出输入源提供的数据
   `fmt`包的`Scan`和`Sscan`开头的函数。代码示例：[readinput.go](/8.IO/readinput.go)
   `Scanln`扫描来自标准输入的文本，将空格分隔的值一次存放到后续的参数内，直到碰到换行。`Scanf`与`Scanln`类似，除了`Scanf`的第一个参数作用格式字符串，用来决定如何读取。以`Sscan`和以`Sscan`开头的函数则是从字符串读取，除此之外，与`Scanf`相同。

## 文件读写

在 Go 语言中，文件使用指向 os.File 类型的指针来表示的，也叫做文件描述符。所以想要通过代码读取一个文件的内容则需要先获取文件描述符。
代码示例：[rwFile.go](/8.IO/rwFile.go)

## 拷贝文件

https://golang.org/pkg/io/#Copy

## 命令行参数处理

在 linux 中使用命令时会用到`SubCommand`和`Options或Flag`，这些子命令和选项，以及命令的帮助信息都是通过`命令行参数处理`这个功能里的各种函数来实现的。该功能也也叫`从命令行读取参数`。并且这些`子命令以及选项或标志`统称为`命令行参数`。备注：linux 中的每一个命令其实就是一个已经编译好的程序。
**flag 包**
使用`flag`包中的相关函数来实现解析命令行的 flag 或 option，详情见：https://golang.org/pkg/flag/#hdr-Usage
下面是其中几种 flag 包的格式说明
格式 1：`flag.TYPE(FlagName, DefaultValue, HelpInfo)`。FlagName 为参数名，DefaultValue 为参数的默认值，HelpInfo 为该参数的帮助信息。返回默认值的指针
格式 2：`flag.TYPE(Pointername ,FlagName, DefaultValue, HelpInfo)`。与上面的格式相同，只不过没有返回值，并且会把`DefaultValue`赋值给`Pointer`指针指向的变量，该变量需要提前定义。
例子：test.go

```
test := flag.String("test","testValue","请指定test参数")
flag.Parse()	//注意必须要有该行才能让test变量获取用户输入的参数，否则一直是默认值
fmt.Println(test)
```

使用方式：`go run test.go -test 123`结果为`123`；若不指定`-test 123`这个参数，则结果为`testValue`。如果使用`go tun test.go -help`则可获得帮助信息`请指定test参数`
**Args**
使用 Go 自带的 Args 切片变量获取命令参数。Args 切片的第一个位置为文件名的绝对路径，第二个位置是使用程序时输入的参数，以空格作为分隔符，分隔每个参数。每个参数都会保存到切片中。e.g.`go run runCommand.go ip`,这时候`ip`的值就会传递到 Args[1]的位置上。
代码示例：[runCommand.go](/8.IO/runCommand.go)和[rwFile.go](/8.IO/rwFile.go)中的函数 7，两个代码里都有从命令行读取参数的实例。

## 启动外部命令和程序

在 Go 程序中，调用外部的系统命令和程序，比如在 Go 程序中执行 linux 命令或者 window 命令。
代码示例：[runCommand.go](/8.IO/runCommand.go)
