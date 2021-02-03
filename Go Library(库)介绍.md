# Go Standard Library(Go标准库) 介绍
[Standard Library(标准库)](https://pkg.go.dev/std) 就是 Go语言 中的 package。每个 package 里有他们对应的常量、变量、函数、方法等。每个库就是一类功能，比如bufio库，这里面就是关于实现读写功能的各种内容；而fmt库则是关于实现格式化输入输出等功能。在[这里](https://pkg.go.dev/std?tab=packages)可以看到 go语言 原生支持的所有标准库。  

与 标准库 相对应的就是 第三方库 ，第三方库一般属于由个人开发，实现更多丰富功能的库。

在[这里](https://pkg.go.dev/)可以搜索自己想要使用的第三方库。

## 标准库首页介绍  
其中`Path`就是包的名称，缩进即是该包下面的子包的名字，右侧有对该包的`Synopsis`i.e.简介。  
## 包页面介绍  
页面中有几个基本的结构：该包的简介、包所包含的各种东西的索引、例子、子包的连接
咱主要需要注意的是这个包里面所包含的东西：一般有如下几种
1. 常量
1. 变量
1. 函数
1. 类型
   1. 函数(处于类型缩进下的函数，其返回值为该类型)
   1. 方法((处于类型缩进下的方法，是由上级类型实现的)
   e.g.os包中的File类型下可以使用Open函数返回File类型，再用Read作用于该类型来实现获取文件中的数据并输出的效果
1. ....每个类型实现的方法都会单独放在该类型下面
注意：有时候函数和方法

# Package
这个 Package 其实就是 Library 概念的一种实例化。
Golang中提供好多自带的`Package`，这些包里有好多已经定义好的常量、变量、类型、函数、方法等可以直接拿来使用。  

**包中各种东西的使用方法**  
在使用的时候，需要在前面加上`包名`和`.`。e.g.我如果要使用os包里的File类型，则格式为：`os.File`，其余的函数方法等一样，都需要在具体的名字前面加上包名。   

# The Core Packages核心包
1. [strings包](#strings包)  
2. [io包i.e.Input/Output](#io包i.e.Input/Output)  
3. Files & Folders包  
4. Errors包  
5. Containers & Sort包  
6. Hashes & Cryptography包  
7. Servers包  
8. Parsing Command Line Arguments包  
9. Synchronization Primitives包  

## strings包  
用于完成对字符串的主要操作。与字符串相关的类型转换都是通过该包来实现的。  
## io包i.e.Input/Output
`io`包由一些函数组成，包里的大多数接口被其它包使用。 其中两个最主要的接口是`Reader`和`Writer`。`Reader`接口通过`Read`方法支持读取操作；`Writer`接口通过`Wrtie`方法支持写入操作。Go语言中的许多函数都将`Reader`和`Writer`这两个方法作为参数。  
## Files&Folders
在Go语言中，使用`os`包中的`Open`函数来打开一个文件。代码示例：[openFile.go](/experiment/1.rwData/rwFile.go)
## Errors包
## Containers & Sort包