# 概述
> 参考：
> - [官网](https://golang.org/)
> - [GitHub 项目](https://github.com/golang/go)
> - [Google 开放源代码](https://cs.opensource.google/go)
> - [go.dev,Tour(Go 语言之旅，通过在线解析器体验 Go 语言的各种特性)](https://go.dev/tour/list)
> - 学习点：
>    - 马哥教育：《Go运维开发架构师》课程
>       - B站：[2020年2月](https://www.bilibili.com/video/BV19a4y1J79s?p=2)，不全
>       - 腾讯视频：[2019年6月](https://ke.qq.com/course/406096#term_id=103076180)
> - 电子书
>    - [Go语言圣经](https://www.k8stech.net/gopl/chapter0/ch0-02/)
> - 牛逼文章
>    - [Go编程模式](https://coolshell.cn/articles/series/go%e7%bc%96%e7%a8%8b%e6%a8%a1%e5%bc%8f)
> - [Golang官方文档英文版](https://golang.org/doc/)
> - [Golang官方文档中文版](http://docscn.studygolang.com/doc/)
> - [Golang标准库英文版](https://golang.org/pkg/)
> - [Golang标准库中文版](https://studygolang.com/pkgdoc)
>   * [标准库文档使用说明](/标准库文档使用说明.md)
> - [格式化占位符说明](/格式化占位符.md)
> - [An Introduction to Programming in Go](https://www.golang-book.com/books/intro)  
> - [Go入门指南、Go语言圣经、Go Web 编程三合一](https://go.wuhaolin.cn/the-way-to-go/eBook/04.7.html)  
> - [Go入门指南](https://github.com/Unknwon/the-way-to-go_ZH_CN)

Go 是一种开源编程语言，可以轻松构建 **simple(简单)**、**reliable(可靠) **和 **efficient(高效) **的软件。

# Hello World
代码：[hello_world/hello_world.go](/hello_world/hello_world.go)
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

```
运行
```shell
$ go run hello_world.go
Hello World
```

# Go 语言 Keywords(关键字)
> 参考：
> - [官方文档,参考-规范-关键字](https://go.dev/ref/spec#Keywords)

Go语言非常简单，只有25个`关键字(Keywords)`可以使用，记住这25个关键字，就掌握了最基本的Go语言用法。这些关键字是 go 语言保留的，不能用作标识符


`关键字`在编程语言中是指该语言的一个功能，比如下文面的 `var`，就是指声明一个变量，`func` 就是定义一个函数等等。
> Note: if-else 算两个关键字所以在这里一共只写了24个。

1. **break** # 控制结构  
1. **case** # 控制结构  
1. **chan** # 用于channel通讯  
1. **const** # 语言基础里面的常量申明  
1. **continue** # 控制结构  
1. **default** # 控制结构  
1. **defer **# 用于在函数退出之前执行某语句的功能  
1. **fallthrough** # 控制结构  
1. **for** # 控制结构  
1. **func **#** **用于定义函数和方法  
1. **go **# 用于并发  
1. **goto** 控制结构  
1. **if-else **#** **控制结构  
1. **import** 用于定义该文件引用某个包  
1. **interface** # 用于定义接口  
1. **map **# 用于声明map类型数据  
1. **package** # 用于定义该文件所属的包  
1. **range** # 用于读取slice、map、channel数据  
1. **return **# 用于从函数返回。有时候也用来直接跳出当前函数，回到主程序继续执行  
1. **select** # 用于选择不同类型的通讯  
1. **struct **# 用于定义抽象数据类型  
1. **switch** # 控制结构  
1. **type **# 用于 Definitions(定义) 自定义类型  或 Declarations(声明) 一个类型的别名。
   1. 其实所谓的类型的别名，也可以当作一种自定义的类型。
1. **var **# 用于 Declarations(声明) 变量 


# Go 语言规范
> 参考：
> - [官方文档,参考-规范](https://go.dev/ref/spec)

干净、可读的代码和简洁性是 Go 追求的主要目标。通过 gofmt 来强制实现统一的代码风格。Go 语言中对象的命名也应该是简洁且有意义的。像 Java 和 Python 中那样使用混合着大小写和下划线的冗长的名称会严重降低代码的可读性。名称不需要指出自己所属的包，因为在调用的时候会使用包名作为限定符。返回某个对象的函数或方法的名称一般都是使用名词，没有 `Get...` 之类的字符，如果是用于修改某个对象，则使用 `SetName`。有必须要的话可以使用大小写混合的方式，如 MixedCaps 或 mixedCaps，而不是使用下划线来分割多个名称。

- go 的命名推荐使用驼峰命名法，必须以一个字母(Unicode字母)或下划线开头，后面可以跟任意数量的字母、数字或下划线。
- 根据首字母的大小写来确定可以访问的权限。无论是方法名、常量、变量名还是结构体的名称，如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用。可以简单的理解成，首字母大写是公有的，首字母小写是私有的
- 不允许循环导入包，即 A 导入 B 包，B 又导入 A 包，这种行为是不可以的。
- 结构体中的属性名需要大写
- 实现结构体的方法，必须与结构体在相同的包中，不允许跨包使用方法来实现结构体。
# 如何阅读Go语言的代码
Go程序的代码从上到下，从左到右阅读，在上例中的代码第一行`package main`称为**package declaration包的声明**。每个Go程序必须以包的声明作为开头第一行。  


一个大的程序是由很多小的基础构件组成的。变量保存值，简单的加法和减法运算被组合成较复杂的表达式。基础数据类型被聚合为数组或结构体等更复杂的数据类型。然后使用if和for之类的控制语句来组织和控制表达式的执行流程。然后多个语句被组织到一个个函数中，以便代码的隔离和复用。函数以源文件和包的方式被组织。   


## 包的概念、导入与可见性
包是结构化代码的一种方式：每个程序都由包（通常简称为 pkg）的概念组成，可以使用自身的包或者从其它包中导入内容。


如同其它一些编程语言中的类库或命名空间的概念，每个 Go 文件都属于且仅属于一个包。一个包可以由许多以 `.go` 为扩展名的源文件组成，因此文件名和包名一般来说都是不相同的。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：`package main`。`package main`表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 `main` 的包。一个应用程序可以包含不同的包，而且即使你只使用 main 包也不必把所有的代码都写在一个巨大的文件里：你可以用一些较小的文件，并且在每个文件非注释的第一行都使用 `package main` 来指明这些文件都属于 main 包。如果你打算编译包名不是为 main 的源文件，如 `pack1`，编译后产生的对象文件将会是 `pack1.a` 而不是可执行程序。另外要注意的是，所有的包名都应该使用小写字母。


**标准库**
在 Go 的安装文件里包含了一些可以直接使用的包，即标准库。在 Windows 下，标准库的位置在 Go 根目录下的子目录 `pkg\windows_386` 中；在 Linux 下，标准库在 Go 根目录下的子目录 `pkg\linux_amd64` 中（如果是安装的是 32 位，则在 `linux_386` 目录中）。一般情况下，标准包会存放在 `$GOROOT/pkg/$GOOS_$GOARCH/` 目录下。


Go 的标准库包含了大量的包（如：fmt 和 os），但是你也可以创建自己的包（第 8 章）。


如果想要构建一个程序，则包和包内的文件都必须以正确的顺序进行编译。包的依赖关系决定了其构建顺序。


属于同一个包的源文件必须全部被一起编译，一个包即是编译时的一个单元，因此根据惯例，每个目录都只包含一个包。


**如果对一个包进行更改或重新编译，所有引用了这个包的客户端程序都必须全部重新编译。**


Go 中的包模型采用了显式依赖关系的机制来达到快速编译的目的，编译器会从后缀名为 `.o` 的对象文件（需要且只需要这个文件）中提取传递依赖类型的信息。如果 `A.go` 依赖 `B.go`，而 `B.go` 又依赖 `C.go`：

- 编译 `C.go`, `B.go`, 然后是 `A.go`.
- 为了编译 `A.go`, 编译器读取的是 `B.o` 而不是 `C.o`.



这种机制对于编译大型的项目时可以显著地提升编译速度。


**每一段代码只会被编译一次**
一个 Go 程序是通过 `import` 关键字将一组包链接在一起。`import "fmt"` 告诉 Go 编译器这个程序需要使用 `fmt` 包（的函数，或其他元素），`fmt` 包实现了格式化 IO（输入/输出）的函数。包名被封闭在半角双引号 `""` 中。如果你打算从已编译的包中导入并加载公开声明的方法，不需要插入已编译包的源代码。


如果需要多个包，它们可以被分别导入：


```
import "fmt"
import "os"
```
或：
```
import "fmt"; import "os"
```


但是还有更短且更优雅的方法（被称为因式分解关键字，该方法同样适用于 const、var 和 type 的声明或定义）：
```
import (
"fmt"
"os"
)
```


它甚至还可以更短的形式，但使用 gofmt 后将会被强制换行：
```
import ("fmt"; "os")
```


当你导入多个包时，最好按照字母顺序排列包名，这样做更加清晰易读。


如果包名不是以 `.` 或 `/` 开头，如 `"fmt"` 或者 `"container/list"`，则 Go 会在全局文件进行查找；如果包名以 `./` 开头，则 Go 会在相对目录中查找；如果包名以 `/` 开头（在 Windows 下也可以这样使用），则会在系统的绝对路径中查找。


导入包即等同于包含了这个包的所有的代码对象。


除了符号 `_`，包中所有代码对象的标识符必须是唯一的，以避免名称冲突。但是相同的标识符可以在不同的包中使用，因为可以使用包名来区分它们。


包通过下面这个被编译器强制执行的规则来决定是否将自身的代码对象暴露给外部文件：


**可见性规则**
当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）。


（大写字母可以使用任何 Unicode 编码的字符，比如希腊文，不仅仅是 ASCII 码中的大写字母）。


因此，在导入一个外部包后，能够且只能够访问该包中导出的对象。


假设在包 pack1 中我们有一个变量或函数叫做 Thing（以 T 开头，所以它能够被导出），那么在当前包中导入 pack1 包，Thing 就可以像面向对象语言那样使用点标记来调用：`pack1.Thing`（pack1 在这里是不可以省略的）。因此包也可以作为命名空间使用，帮助避免命名冲突（名称冲突）：两个包中的同名变量的区别在于他们的包名，例如 `pack1.Thing` 和 `pack2.Thing`。你可以通过使用包的别名来解决包名之间的名称冲突，或者说根据你的个人喜好对包名进行重新设置，如：`import fm "fmt"`。下面的代码展示了如何使用包的别名：


示例
```
package main
import fm "fmt" // alias3
func main() {
   fm.Println("hello, world")
}
```


**注意事项**
如果你导入了一个包却没有使用它，则会在构建程序时引发错误，如 `imported and not used: os`，这正是遵循了 Go 的格言：“没有不必要的代码！“。


**包的分级声明和初始化**
你可以在使用 `import` 导入包之后定义或声明 0 个或多个常量（const）、变量（var）和类型（type），这些对象的作用域都是全局的（在本包范围内），所以可以被本包中所有的函数调用，然后声明一个或多个函数（func）。


## 函数
这是定义一个函数最简单的格式：
```go
func FunctionName()
```
你可以在括号 `()` 中写入 0 个或多个函数的参数（使用逗号 `,` 分隔），每个参数的名称后面必须紧跟着该参数的类型。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。如果你的 main 包的源代码没有包含 main 函数，则会引发构建错误 `undefined: main.main`。main 函数既没有参数，也没有返回类型（与 C 家族中的其它语言恰好相反）。如果你不小心为 main 函数添加了参数或者返回类型，将会引发构建错误： 


   func main must have no arguments and no return values results.


在程序开始执行并完成初始化后，第一个调用（程序的入口点）的函数是 `main.main()`（如：C 语言），该函数一旦返回就表示程序已成功执行并立即退出。函数里的代码（函数体）使用大括号 `{}` 括起来。左大括号 `{` 必须与方法的声明放在同一行，这是编译器的强制规定，否则你在使用 gofmt 时就会出现错误提示：   `build-error: syntax error: unexpected semicolon or newline before {`（这是因为编译器会产生 `func main() ;` 这样的结果，很明显这错误的）


**Go 语言虽然看起来不使用分号作为语句的结束，但实际上这一过程是由编译器自动完成，因此才会引发像上面这样的错误**


右大括号 `}` 需要被放在紧接着函数体的下一行。如果你的函数非常简短，你也可以将它们放在同一行：


```
func Sum(a, b int) int { return a + b }
```


对于大括号 `{}` 的使用规则在任何时候都是相同的（如：if 语句等）。


因此符合规范的函数一般写成如下的形式：
```
func functionName(parameter_list) (return_value_list) {
   …
}
```
其中：

- parameter_list 的形式为 (param1 type1, param2 type2, …)
- return_value_list 的形式为 (ret1 type1, ret2 type2, …)



只有当某个函数需要被外部包调用的时候才使用大写字母开头，并遵循 Pascal 命名法；否则就遵循骆驼命名法，即第一个单词的首字母小写，其余单词的首字母大写。


下面这一行调用了 `fmt` 包中的 `Println` 函数，可以将字符串输出到控制台，并在最后自动增加换行字符 `\n`：


```
fmt.Println（"hello, world"）
```


使用 `fmt.Print("hello, world\n")` 可以得到相同的结果。`Print` 和 `Println` 这两个函数也支持使用变量，如：`fmt.Println(arr)`。如果没有特别指定，它们会以默认的打印格式将变量 `arr` 输出到控制台。单纯地打印一个字符串或变量甚至可以使用预定义的方法来实现，如：`print`、`println：print("ABC")`、`println("ABC")`、`println(i)`（带一个变量 i）。这些函数只可以用于调试阶段，在部署程序的时候务必将它们替换成 `fmt` 中的相关函数。当被调用函数的代码执行到结束符 `}` 或返回语句时就会返回，然后程序继续执行调用该函数之后的代码。程序正常退出的代码为 0 即 `Program exited with code 0`；如果程序因为异常而被终止，则会返回非零值，如：1。这个数值可以用来测试是否成功执行一个程序。


## 注释
示例 
```
package main
import "fmt" // Package implementing formatted I/O.
func main() {
   fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
}
```


上面这个例子通过打印 `Καλημέρα κόσμε; or こんにちは 世界` 展示了如何在 Go 中使用国际化字符，以及如何使用注释。


注释不会被编译，但可以通过 godoc 来使用（第 3.6 节）。


单行注释是最常见的注释形式，你可以在任何地方使用以 `//` 开头的单行注释。多行注释也叫块注释，均已以 `/*` 开头，并以 `*/` 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。


每一个包应该有相关注释，在 package 语句之前的块注释将被默认认为是这个包的文档说明，其中应该提供一些相关信息并对整体功能做简要的介绍。一个包可以分散在多个文件中，但是只需要在其中一个进行注释说明即可。当开发人员需要了解包的一些情况时，自然会用 godoc 来显示包的文档说明，在首行的简要注释之后可以用成段的注释来进行更详细的说明，而不必拥挤在一起。另外，在多段注释之间应以空行分隔加以区分。


示例：
```go
// Package superman implements methods for saving the world.
//
// Experience has shown that a small number of procedures can prove
// helpful when attempting to save the world.
package superman
```


几乎所有全局作用域的类型、常量、变量、函数和被导出的对象都应该有一个合理的注释。如果这种注释（称为文档注释）出现在函数前面，例如函数 Abcd，则要以 `"Abcd..."` 作为开头。


示例：


```
// enterOrbit causes Superman to fly into low Earth orbit, a position
// that presents several possibilities for planet salvation.
func enterOrbit() error {
...
}
```


godoc 工具（第 3.6 节）会收集这些注释并产生一个技术文档。


## 类型


变量（或常量）包含数据，这些数据可以有不同的数据类型，简称类型。使用 var 声明的变量的值会自动初始化为该类型的零值。类型定义了某个变量的值的集合与可对其进行操作的集合。


类型可以是基本类型，如：int、float、bool、string；结构化的（复合的），如：struct、array、slice、map、channel；只描述类型的行为的，如：interface。


结构化的类型没有真正的值，它使用 nil 作为默认值（在 Objective-C 中是 nil，在 Java 中是 null，在 C 和 C++ 中是NULL或 0）。值得注意的是，Go 语言中不存在类型继承。


函数也可以是一个确定的类型，就是以函数作为返回类型。这种类型的声明要写在函数名和可选的参数列表之后，例如：


```
func FunctionName (a typea, b typeb) typeFunc
```


你可以在函数体中的某处返回使用类型为 typeFunc 的变量 var：
```
return var
```


一个函数可以拥有多返回值，返回类型之间需要使用逗号分割，并使用小括号 `()` 将它们括起来，如：
```
func FunctionName (a typea, b typeb) (t1 type1, t2 type2)
```


示例： 函数 Atoi (第 4.7 节)：`func Atoi(s string) (i int, err error)`


返回的形式：
```
return var1, var2
```




这种多返回值一般用于判断某个函数是否执行成功（true/false）或与其它返回值一同返回错误消息（详见之后的并行赋值）。


使用 type 关键字可以定义你自己的类型，你可能想要定义一个结构体(第 10 章)，但是也可以定义一个已经存在的类型的别名，如：


```
type IZ int
```


**这里并不是真正意义上的别名，因为使用这种方法定义之后的类型可以拥有更多的特性，且在类型转换时必须显式转换。**


然后我们可以使用下面的方式声明变量：


```
var a IZ = 5
```


这里我们可以看到 int 是变量 a 的底层类型，这也使得它们之间存在相互转换的可能（第 4.2.6 节）。


如果你有多个类型需要定义，可以使用因式分解关键字的方式，例如：


```
type (
   IZ int
   FZ float64
   STR string
)
```


每个值都必须在经过编译后属于某个类型（编译器必须能够推断出所有值的类型），因为 Go 语言是一种静态类型语言。


## Go 程序的一般结构


下面的程序可以被顺利编译但什么都做不了，不过这很好地展示了一个 Go 程序的首选结构。这种结构并没有被强制要求，编译器也不关心 main 函数在前还是变量的声明在前，但使用统一的结构能够在从上至下阅读 Go 代码时有更好的体验。


所有的结构将在这一章或接下来的章节中进一步地解释说明，但总体思路如下：


- 在完成包的 import 之后，开始对常量、变量和类型的定义或声明。
- 如果存在 init 函数的话，则对该函数进行定义（这是一个特殊的函数，每个含有该函数的包都会首先执行这个函数）。
- 如果当前包是 main 包，则定义 main 函数。
- 然后定义其余的函数，首先是类型的方法，接着是按照 main 函数中先后调用的顺序来定义相关函数，如果有很多函数，则可以按照字母顺序来进行排序。



示例  


```go
package main

import (
   "fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
}

func main() {
   var a int
   Func1()
   // ...
   fmt.Println(a)
}

func (t T) Method1() {
   //...
}

func Func1() { // exported function Func1
   //...
}
```


Go 程序的执行（程序启动）顺序如下：

1. 按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
1. 如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
1. 然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
1. 在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。



## 类型转换


在必要以及可行的情况下，一个类型的值可以被转换成另一种类型的值。由于 Go 语言不存在隐式类型转换，因此所有的转换都必须显式说明，就像调用一个函数一样（类型在这里的作用可以看作是一种函数）：
```
valueOfTypeB = typeB(valueOfTypeA)
```


**类型 B 的值 = 类型 B(类型 A 的值)**


示例： 


```
a := 5.0
b := int(a)
```


但这只能在定义正确的情况下转换成功，例如从一个取值范围较小的类型转换到一个取值范围较大的类型（例如将 int16 转换为 int32）。当从一个取值范围较大的转换到取值范围较小的类型时（例如将 int32 转换为 int16 或将 float32 转换为 int），会发生精度丢失（截断）的情况。当编译器捕捉到非法的类型转换时会引发编译时错误，否则将引发运行时错误。


具有相同底层类型的变量之间可以相互转换：


```
var a IZ = 5
c := int(a)
d := IZ(c)
```




# [gopher](https://www.cnblogs.com/laud/p/gopher.html)
**原文链接：**[**https://www.cnblogs.com/laud/p/gopher.html**](https://www.cnblogs.com/laud/p/gopher.html)
**​**

#### 前言
gopher原意地鼠，在golang 的世界里解释为**地道的go程序员**。在其他语言的世界里也有PHPer，Pythonic的说法，反而Java是个例外。虽然也有Javaer之类的说法，但似乎并不被认可。而**地道**或者说**道地**，说的是gopher写的代码无不透露出go的独特气息，比如项目结构、命名方式、代码格式、编码风格、构建方式等等。用gopher的话说，用go编写代码就像是在画一幅中国山水画，成品美不胜收，心旷神怡。
#### 环境变量
gopher第一条：把东西放对地方。
go程序的运行，需要依赖于两个基础的环境变量，GOROOT与GOPATH。环境变量几乎在各类编程语言中都存在，比如java的JAVA_HOME，其实也就是编译器及相关工具或标准库所在目录。但go除了GOROOT之外，还增加了GOPATH，它指的是go程序依赖的第三方库或自有库所在目录，以指示编译器从这些地方找到依赖。GOPATH支持多个目录，通常一个目录就是一个项目，并且GOPATH目录按约定由src、pkg、bin三个目录组成。gopher们的做法是定义Global GOPATH、Project GOPATH，而更大的项目还会定义Module GOPATH。当使用go get下载依赖时，会选择GOPATH环境变量中的第一个目录存放依赖包。

| 变量 | 含义 | 说明 |
| --- | --- | --- |
| GOROOT | go运行环境根目录 | 通常指go sdk安装目录，包含编译器、官方工具及标准库 |
| GOPATH | 工作环境目录列表 | 通常指第三方库、自有库及项目目录 |

##### GOPATH

1. src: 源代码目录
1. pkg: 编译时中间文件所在目录
1. bin: 编译后生成可执行文件所在目录
```
├── src
|   └── eventer.com
|       └── test
|           └── main.go
├── pkg
└── bin
|   └── test.bin
```
##### Global GOPATH
全局GOPATH，通常将第三方库的代码保存至此处，便于所有项目引用，而不用重复下载。因此在GOPATH中，建议将global gopath放在最前面。
##### Project GOPATH
如果项目比较简单的话，可以采用global GOPATH与project GOPATH的组合，即GOPATH设置为global GOPATH与project OGPATH两个目录。比如引用了需要对源码稍作修改的开源项目时，开源项目可以跟当前项目一起放在src目录下。
##### Module GOPATH
在比较复杂的项目中，如果划分的模块比较多，则可以采用global, project, module三种GOPATH组合的方式，即即GOPATH设置为global GOPATH、project OGPATH、module GOPATH三个目录。
最后请注意，global GOPATH、project GOPATH、module GOPATH不是go语言中的概念，go要求的只有GOPATH。而因为GOPATH是一个目录列表，所以我只是在实践中将GOPATH细分为上述三者，同时便于讲述GOPATH的概念。
#### 项目结构
gopher第二条：按东西放在约定的地方。
不同复杂度的项目，大致可以有两种类型的项目结构，第一种是依赖较少、较简单的项目，如下图所示，其中mydriver项目是假设要引用的开源项目，但要在当前项目的使用中稍修改，因此与当前项目的源码放在一起。
```
├── src
|   ├── github.com
|       └── eventer
|           └── mydriver
|               └── driver.go
|   └── eventer.com
|       └── test
|           └── main.go
├── pkg
└── bin
```
第二种是依赖较多、模块划分较细、较复杂的项目，如下图所示：
```
├── src
|   └── eventer.com
|           └── test
|               └── main.go
├── pkg
├── bin
├── module1
|   ├── src
|       └── eventer.com
|           └── lib
|               └── lib1.go
|   ├── pkg
|   └── bin
├── module2
|   ├── src
|       └── eventer.com
|           └── lib
|               └── lib2.go
|   ├── pkg
|   └── bin
└── module3
|   ├── src
|       └── eventer.com
|           └── lib
|               └── lib3.go
|   ├── pkg
|   └── bin
```
其中的module1，module2，module3是基础模块，被主项目依赖。这个时候除了将主项目路径放在GOPATH外，还应当将基础模块路径放在GOPATH中。这样的项目结构不仅让项目清晰，还可以做不同的分工。

#### 命名规范
gopher第三条：把名字起得go一点。
go语言的命名与其他语言最大的不同在于首字母的大小写。大写代表公开（导出，可以在其他包内访问）；小写代表私有（不导出，只能在包内访问）。除此之外，与其他语言并无二致，比如不能以数字开头。而由于关键字、保留字的减少，因而减少了一些命名上的忌讳。更为突出的是，go语言有一些建议性的命名规范，这也是gophers的圣经，理应严格遵守。

| 约定 | 范围 | 说明 | 示例 |
| --- | --- | --- | --- |
| 驼峰命名法 | 全局 | 统一使用驼峰命名法 | var isLive = false |
| 大小写一致 | 缩写短语，惯用词 | 如HTML，CSS, HTTP等 | htmlEscape，HTMLEscape |
| 简短命名法 | 局部变量 | 方法内、循环等使用的局部变量可以使用简短命名 | 比如for循环中的i，buf代表buffer等 |
| 参数命名法 | 函数参数、返回值、方法接收者 | 如果参数类型能说明含义，则参数名可以简短，否则应该采用有文档说明能力的命名 | 比如d Duration，t Time |
| 通用命名法 | 作用域越大或者使用的地方离声明的地方太远，则应采用清晰有意义的命名 | - |  |
| 导出命名法 | 导出变量、函数、结构等 | 包名与导出名意义不要重复，同时包的命名要与导出的内容相关，不要使用宽泛的名字，如common，util | bytes.Buffer比bytes.ByteBuffer要好 |
| 文件命名 | go文件，单元测试文件 | go文件名尽量以一个单词来命名，多个单词使用下线划分隔，单元测试文件以对应go文件名加_test结尾 | proto_test |
| 包命名 | 包 | 包的一级名称应是顶级域名，二级名称则应是项目名称，项目名称单词间以-分隔 | github.com/mysql |

#### 代码格式
gopher第四条：按统一的格式来。
在多人协作团队中，统一的代码格式化模板是第一要义。在Java语言中，检验新人经验的一大法宝就是他有没有主动索要代码模板。而在go语言中，则没有这个必要了。因为go已经有默认的代码格式化工具了，而且代码格式化在go语言中是强制规范。所以这使得所有go程序员写出来的代码格式都是一样的。
go默认的代码格式化工具是gofmt。另外还有一个增强工具goimport，在gofmt的基础上增加了自动删除和引入依赖包。而行长则以不超过80个字符为佳，超过请主动以多行展示。
#### 编码风格
gopher第五条：请学会约定

| 项 | 约定 | 说明 |
| --- | --- | --- |
| import | 按标准库、内部包、第三方包的顺序导入包 | 只引一个包时使用单行模式，否则使用多行模式 |
| 变量声明 | 如果连续声明多个变量，应放在一起 | 参见例子 |
| 错误处理 | 不要忽略每一个error，即使只是打一行日志 | go的error处理方式与C同出一辙，通过返回值来标明错误或异常，引来的争议也很多，甚至官方已经开始酝酿在go2解决这个问题 |
| 长语句打印 | 使用格式化方式打印 | - |
| 注释规范 | 变量、方法、结构等的注释直接加上声明前，并且不要加空行。废弃方法加Deprecated:即可 | 其中的第一行注释会被godoc识别为简短介绍，第二行开始则被认为是注释详情。注释对godoc的生成至关重要，因此关于注释会有一些技巧，我将在后面用专门的章节探讨 |

多变量声明
```
var (
    name string
    age int
)
```
注释规范
```
// Add 两数相加
// 两个整数相加，并返回和。
func Add(n1, n2 int)int{
    return n1 + n2
}
```
#### 依赖管理
gopher第六条：使用依赖管理工具管理自有依赖与第三方依赖
一个语言的生态是否完善甚至是否强大，除了github上面的开源项目数量之外，还有一大特征就是是否有优秀的依赖管理工具。依赖管理工具在业界已经是无处不在，yum、maven、gradle、pip、npm、cargo这些工具的大名如雷贯耳。那么go有什么呢？
早期go的依赖是混乱的，因为没有一个工具能得到普遍认可，而官方又迟迟不出来解决问题。历数存在的工具包括godep、glide、govender等等。甚至早期还需要使用GOPATH来管理依赖，即项目的所有依赖都通过go get下载到指定的GOPATH中去。当然这种方案还可以撑大多数时间，但随着时间的流逝，随着开发人员的变动，这种管理依赖的弊端就慢慢显现出来。其实这些老路早期的java也走过，曾几何时，每个java项目里面都会有一个叫lib或libs的目录，这里放的就是当前项目依赖的包。当GO采用GOPATH来管理依赖时，开发人员只能被倒逼着用java的方式在源码库中自行管理依赖。这样相当于给依赖包做了隔离，同时又具备了版本管理（因为放在源码库）。
后来在go1.5的时候，官方引入了vender的概念，其实这也没改变多少，只是官方让大家存放依赖包的目录名称不要乱起了，统一叫vender吧。这个方案我觉得比依赖GOPATH还糟糕，因为vendor目录脱离了版本管理，导致更换依赖包版本很困难，在当前项目对依赖包的版本更新可能会影响其他项目的使用（如果新版本的依赖包有较大变动的话），同时如何将依赖包放到vendor下呢？等等。当然官方做出的这些变动可能是想像maven那样，推动社区来完成这件事，因而直接推动了上文提到的基于vendor的依赖管理工具的诞生。直至后来官方默认的社区做出来dep，这下安静了，尽管刚开始时也不怎么好用，但有总比没有好。
go1.11在vgo的基础上，官方推出了go module。在发布前，官方与社区的大神们还为此开吵，认为官方太不厚道且独断专行。完全忽视dep社区的存在，无视dep在go语言中的地位与贡献。喜欢八卦的朋友们，可搜索《关于Go Module的争吵》一览大神是怎么吵架的，也可从中学习他们的思想。
作为dep与module的亲身实践者，还是乖乖地用dep吧。除非你有足够的热情去折腾，比如弄个高速的科学上网工具，或者搭建Go module proxy。开源的go module proxy比如athens，goproxy。《Hello，Go module proxy》一文详细介绍了go module的幸福与烦恼，反正我是没有感到幸福。此文很全面地介绍go module后遗症，goproxy可以解决科学上网的问题，然后也有了athens这样的工具出现，go的依赖管理是朝着越来越好的方向发展。dep与module的争议在于respect and free，至少目前看来，两者是可以共存的，特别是在国内。
相对于java的依赖管理工具maven或gradle来说，gradle是maven的升级版，同时带来了DSL与元编程的特性，这无疑使得gradle异常地强大。但gradle.io在国内的可达情况也不尽如人意，好就好在其与maven仓库标准的兼容，使得从maven转到gradle几乎没有额外的成本及阻力。
扯了这么多，依赖管理对于一门语言是必不可少的。c有cmake，java有maven、gradle，rust有cargo，那么go的dep或者module就用起来吧，看完大神吵架之后，喜欢哪个就选哪个。是不可能产生一个能满足所有人要求的依赖管理工具的，就连号称最牛逼的cargo也不例外。在一般的项目中，能用到的依赖管理功能也就那常用的几个而已，对大多数项目来说，适用好用就行。
#### 构建方式
gopher第七条：按需构建
构建的目标是让代码成为可运行程序。构建的过程应该是低成本并且让人愉悦的，显然C在这一方面让人抓狂，而go确实做得不错。并且能在任何平台下编译出另外一个平台的可执行程序。不管你的go程序是CLI、GUI、WEB或者其他形式的网络通讯程序，在go的世界里都只需要一个命令构建成可执行程序（依赖也一并被打包），即可在目标系统上运行。在这一点上，java是望尘莫及了。
下面是用来构建go程序常用的参数，其他参数可通过go help environment命令查看。

| 参数 | 值 | 说明 |
| --- | --- | --- |
| CGO_ENABLED | 0 or 1 | 是否支持cgo命令，如果go代码中有c代码，需要设置为1 |
| GOOS | darwin, freebsd, linux, windows | 可执行程序运行的目标操作系统 |
| GOARCH | 386, amd64, arm | 可执行程序运行的目标操作系统架构 |

```
# Linux下编译Mac 64位可执行程序
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
# Linux下编译windows 64位可执行程序
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
# 如果想减少二进制程序的大小，可以加上-ldflags "-s -w"，但同时会丢掉调试信息，即不能用gdb调试了。
# 如果想更进一步减少程序大小，可以使用加壳工具，比如upx
```

# 版本变化
Go 1.16 将会弃用 io/ioutil 包，ioutil 包中的功能将会由 io 及 os 包所替代。详见：[#40025](https://github.com/golang/go/issues/40025)、[Go 1.16 Release Notes](https://golang.org/doc/go1.16#ioutil)

## 名词缩写
###	A
ARG：Arguments参数。多指实际参数、实参  
ASYNC：asynchronization异步。  
###	B
bool：boolean布尔、布尔型、布尔值  
###	C
CAP：capacity容量。 
chan：channel通道。
CSP：Communicating Sequential Processes顺序通信处理。  
CLI：Command Line Interface命令行接口。常代指使用命令来对程序进行操作，相对的非命令行是GUI
###	D
DEV：development开发。常指软件开发人员
DOC：documentation文档。常用来代指某项技术的使用说明书
###	E
EXP：expression表达式。  
EOF：End Of  File文件的末尾。是操作系统无法从数据源读取更多数据的情形。  
###	F
FD:File Descriptor文件描述符。  
FIELD：field字段。  
###	G
GFS：Distributed File System分布式文件系统
GUI：Graphical User Interface图形用户界面。常指通过图形界面来对程序进行操作，相对的是非图形CLI
###	H
###	I
ID：identifier标识符。一般用于指各个功能的名字e.g.变量名、map名、数组名、切片名、函数名等等。  
IDE：Integrated Development Environment集成开发环境。
init：initialization初始化。
int：integers整数、整数型。  
###	J
###	K
KEY/VAL：key/value键/值。  
###	L
LEN：length长度。  
###	M
###	N
###	O
OPS：operations运维。常代指IT运维技术人员
###	P
PARAM：parameter参数。多值形式参数、形参。  
PANIC：panic恐慌。一般代指严重错误。  
###	Q
###	R
RECV：receive接收。  
###	S
SEND：send发送。  
SPEC: specification说明书、属性说明、规格
SYNC：synchronization同步。  
###	T
###	U
###	V
VAR：variable变量  
###	W
###	X
###	Y
###	Z


