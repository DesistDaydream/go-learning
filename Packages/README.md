# Package  
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
























