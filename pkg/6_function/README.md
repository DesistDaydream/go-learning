# Functions函数
1. Your Second Function第二个函数
2. Returning Multiple Values返回多个值
3. [Variadic Functions](#Variadic_Functions) 可变参数函数
4. [Closure](#Closure) 闭包
5. [Recursion](#Recursion)递归i.e.指递归函数
6. [Other](#Other)
   1. [make()与new()](#make()与new())

DRY原则：Don't Repeart Yourself不要重复你自己
Golang有3中类型的函数：
   1. 普通的带有名字的函数
   2. 匿名函数或者lambda函数
   3. 方法(Methods)
除了main()、init()函数外，其余所有类型的函数都可以有参数与返回值。函数参数、返回值以及它们的类型统称为**函数签名**。

## 定义函数
```
func FunctionName([Parameter]) [(returnValue)] {
	代码块
}
```
   1. 其中`()`里的Parameter以及returnValue可以省略，但是至少要包含一个`()`，哪怕这个小括号内没有任何内容。i.e.一个函数可以没有参数，与返回值，仅仅执行本身所提供的功能
   2. `Parameter`形式参数。这是一个参数列表，包括参数名以及参数类型。参数一般情况是变量、或者另一个函数(这个函数也可以当做变量来使用，是函数类型的变量，在调用时，可以通过实参改变该函数)。
   3. `returnValue`返回值。同样包括参数名以及参数类型，参数一般是变量。可以直接定义变量名与类型，也可以省略变量名直接指明返回值的类型
   4. 其中 Parameter 与 returnValue 都是可省的，最简化的定义格式为 `function Name(){}`

## 调用函数i.e.使用函数
格式为：
`[Pack.]Function([ARG1, ARG2, ..., ARGn])`
其中 Pack 与 ARG 都是可省的，若在同一个包中，则 Pack 可省，若不用传递参数，则 ARG 可省。
包名.函数名(实际参数)。`Function`是`Pack`包里的一个函数，括号里的是被调用函数的实参，这些实参的值被传递给被调用函数的形参，参数的传递是一一对应的，第一位传递给第一位，第二位传递给第二位，以此类推。在引用的时候参数可省略为空，但是括号必须要有
   1. 实参：actual parameter实际参数,一般用arguments表示，在调用函数时使用实参
   2. 形参：formal parameter形式参数，一般用parameter表示，在定义函数时使用形参
注意：
```
在使用函数返回值对变量进行赋值的时候，可以使用`_`下划线，来把函数的某一个返回值丢弃。详见[return.go](/5.Functions/return.go)中下划线的使用方法
如果在同一个包下的多个文件之间互相调用函数，在执行 go run XXX.go 命令时，需要指定所有文件，即 go run *.go 只有这样在引用其他文件中的函数时，才可以成功。
```

### 回调函数
函数可以作为其他函数的参数进行传递，然后在其它函数内调用执行，一般称之为**回调**。e.g.[function_parameter.go](/5.Functions/function_parameter.go)

### 内置函数
一些不需要进行导入操作就可以使用的函数就是内置函数。
通俗来说，是在不使用`import`关键字导入任何包，就可以在main下直接使用的各种函数。

|名称|说明|
|----------------|----------------|
|close|用于管道通信|
|len、cap|len用于返回某个类型的长度或数量(字符串、数组、切片、map和管道)；cap是容量的意思，用于返回某个类型的最大容量(只能用于切片和map)|
|new、make|new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：v := new(int)。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作`new()`是一个函数，不要忘记它的括号|
|copy、append|用于复制和连接切片|
|panic、recover|两者均用于错误处理机制|
|print、println|底层打印函数，在部署环境中建议使用fmt包内的打印函数|
|complex、real imag|用于创建和操作复数|

# 函数的其余功能

## Variadic_Functions
代码示例[variadic.go](/5.Functions/variadic/variadic.go)
如果函数的最后一个参数是采用`...type`的形式，那么这个函数就可以处理一个变长的参数，这个长度可以为0，这样的函数称为VariadicFunctions。
格式为：`func Funcition(Param1,Param2,...ParamX ...TYPE) {}`。最后一个参数`ParamX`类型为`TYPE切片`，这个参数称为`变长参数`，该参数可接收多个值。
e.g. 示例函数以及调用
```
func Greeting(prefix string, who ...string)
Greeting("hello:", "Joe", "Anna", "Eileen")
```
该例子中，变量`who`的值为`[]string{"Joe", "Anna", "Eileen"}`

这类函数可以接受**任意数量的形参**，因为使用的的`...`作为形参列表

## Closure
**Closure(闭包)** 通俗得讲，就是函数 a 的内部函数 b，被函数 a 外部的一个变量引用的时候，就创建了一个闭包代码示例 [closure.go](/5.Functions/closure.go)

**Anonymous Functions(匿名函数)**，当我们不希望给函起名字的时候可以使用`匿名函数`。

匿名函数不能够独立存在：
- 可以把匿名函数赋值于某个变量。i.e.保存函数的地址到变量中。
   - `fplus := func(x, y int) int { return x+y }`
      - 然后可以同通过变量名对函数进行调用i.e.`fplus(3,4)`
- 可以让匿名函数仅仅实现自身逻辑即可
   - `func(x, y int) { fmt.Println(x + y) }(3, 4)`
      - 第一个括号 `(x, y int)` 表示该函数的形参
      - 最后的括号 `(3, 4)` 表示调用该函数的实参

**匿名函数**可以构造**闭包Closure**。

闭包的特性：
1. 封闭性：外界无法访问闭包内部的数据，如果在闭包内声明变量，外界是无法访问的，除非闭包主动向外界提供范文接口
2. 持久性：一般的函数，在调用完毕后，系统自动注销函数，而对于闭包来说，在外部函数被调用之后，闭包的结构依然保存在系统中，闭包中的数据依然存在
3. 闭包会占用内存资源，过多的使用会导致内存溢出。

闭包是一个有状态(不消失的私有数据)的函数。闭包是一个有记忆的函数。闭包相当于一个只有一个方法的紧凑对象。

## Recursion
当一个函数在其函数体内调用自身，则称之为**递归**。最经典的例子是计算斐波那契数列，即前两个数为1，从第三个数开始每个数均为前两个数之和。

## Other
展示一个计算过程所消耗的时间：使用time包中的`Now()`和`Sub()`函数，在计算开始之前设置一个起始时间，再由计算结束时的结束时间与起始时间相减，即为执行计算的消耗时间。代码示例:[fibonnaci.go](/5.Functions/fibonacci/fibonacci.go)

### make()与new()

**make()**
make用于 内建类型(map、slice、channel) 的内存分配

`make(T,args)` 与 `new(T)` 有着不同的功能，make 只能创建 slice,map,channel，并且返回一个有初始值(非零)的T类型，而不是`*T`

**new()**
`new()` 用于为*各种类型*分配内存。 第一个参数是类型，而不是值，返回的值是指向该类型新分配的零值的指针。

new(T)分配了零值填充的 `T类型` 的内存空间，并且返回其地址，即一个 `*T` 类型的值(GO语言的术语：返回了一个指针，指向新分配的类型T的零值)

new() 是 go 语言中很重要的一种 **设计思想**,在原生基础包里，仅仅返回一个类型的指针。而在实际项目中，人们常常将 new() 思想与 struct 相结合。比如下面这个例子：
```go
// 声明一个结构体
type MemorySession struct {
	sessionID string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

// NewMemorySession 返回一个储存 Session 的内存引擎
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{}
	return s
}
```
这个例子中，就使用了 new() 的思想，通过一个函数，来生成一个结构体的指针，这个指针通常称为该结构体的 **instance(实例)**。这样，后续要操作这个结构体属性的值，调用这些指针(i.e.实例)即可
```go
s := NeNewMemorySession("DD")
s.METHOD1
s.METHOD2
....等等
```

再比如在 [kubernetes](https://github.com/kubernetes/kubernetes) 项目中，也可以在很多地方看到这类**设计思想**
这里用 kubelet 的[一段代码](https://github.com/kubernetes/kubernetes/blob/2b23e71f6c88a08b9ab43523a86e34c4ba0601ae/pkg/kubemark/hollow_kubelet.go#L65)举例：
```go
type HollowKubelet struct {
   KubeletFlags         *options.KubeletFlags
   KubeletConfiguration *kubeletconfig.KubeletConfiguration
   KubeletDeps          *kubelet.Dependencies
}

func NewHollowKubelet(
   flags *options.KubeletFlags,
 .....

   return &HollowKubelet{
      KubeletFlags:         flags,
      KubeletConfiguration: config,
      KubeletDeps:          d,
   }
}
```
这里用 NewHollowKubelet() 函数实例化了一个 HollowKubelet 结构体。后续想要使用时，调用 New 函数，生成一个实例(指针变量)，然后调用作用在这个结构体上的方法即可。


**总结：**
`new`负责分配内存，`new(T)`返回`*T`指向一个零值`T`的指针
`make` 负责初始化值，`make(T)` 返回初始化后的 `T` ，而非指针
最重要的一点：make 仅适用于slice，map 和channel
关于，并非是空值，而是一种“变量未填充前”的默认值，通常为0，如下
```
int  0
int8 0
int32 0
int64 0
uint 0x0
rune 0 //rune的实际类型是 int32
byte 0x0 //byte的实际类型是uint8
float32 0 //长度为4 byte
float64 0 //长度为8 byte
bool false
string ""
```

### 何时使用new()和make()
* 切片、映射和通道，使用make()
* 数组、结构体和所有的值类型，使用new()