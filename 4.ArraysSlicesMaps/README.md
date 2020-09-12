# Arrays, Slices and Maps数组、切片、映射
1. [Arrays](#Arrays数组)  
2. [Slices](#Slices切片)
3. [Maps](#Maps映射)

## Arrays数组  
数组是具有相同的唯一类型的一组已编号、且长度固定的数据项序列。每个数据项称为**element(元素)**、长度指的是**元素的个数**、编号指每个元素的**index(索引)**从0开始。代码示例：[arrays.go](/4.ArraysSlicesMaps/arrays.go)  
1. 数组的声明。格式为：`var MapID [LENGTH]TYPE`,MapID是标识符i.e.数组的名字，LENGTH是数组长度i.e.元素的个数，TYPE是每个元素的数据的类型。`[]`中括号是数组类型的标识符，不要忘记写。  
   1. e.g.`var arr [5]int`这定义了一个名为arr、长度为5、数据类型为整型的数组。  
2. 数组的初始化。数组定义后，默认初始化每个元素的值为0，后续可以对每个元素进行赋值。数组可以有两种初始化方式
   1. 每次对一个元素进行赋值，一般使用循环来实现
   2. 使用`{}`大括号，直接对数组进行初始化
      1. e.g.`var arr = [5]int{1,5,23,2,10}`
3. 数组的引用。  
   1. 引用数组的长度。格式:`len(MapID)`(使用len()函数，括号内为数组标识符)。也就是元素的数目，必须是固定的并且在声明该数组时就给出，数组长度最大为2Gb。格式为`len(ARRAYS)`len是length的缩写，ARRAYS里使用数组的名称。  
      1. e.g.对于上面例子中定义的数组，数组的长度为`len(arr)`  
  2. 引用数组中的元素。格式：`MapID[INDEX]`(数组标识符[元素编号])。数组的元素可以通过**索引**(数组的位置，索引有时候也叫作数组中元素的**编号**)来读取(或修改)，索引从0开始，第一个元素的索引为0，第二个索引为1，以此类推(长度为3的数组，元素的索引为0、1、2)。  
      1. e.g.对于上面例子中定义的数组，第一个元素是`arr[0]`,第二个元素是`arr[1]`.....第五个元素是`arr[4]`  

### 多维数组
数组通常是一维的，但是可以用来组装成多维数组。e.g.`[3][5]int`有行有列,`[2][2][2]float64`立体效果。代码示例：[multidim_array.go](/4.ArraysSlicesMaps/multidim_array.go)  

## Slices切片  
切片是一个**长度可变的数组**，是**数组的一部分**;是对数组一个连续片段的**引用**。这个片段可以是整个数组、或是由起始和终止索引标识中间的元素子集，注意：终止索引标识符的元素不包括在切片内。代码示例：[slices.go](/4.ArraysSlicesMaps/slices.go)  
切片的由来：Golang中的数组是一个值，数组变量表示整个数组，而不是指向数组第一个元素的指针。这就意味，将一个数组当作参数传递时，会完全拷贝数组中的内容(当数组非常大的时候，会非常占用资源，使用起来也不便利)，这时候就可以是使用切片来作为参数进行传递。**可以把数组当成一个存储元素的地方，具有索引，有着固定的大小。而切片则是指向这个存储元素地方的指针。**所以Golang中一般使用切片来对数组进行引用和传递参数。  
注意：绝对不要用指针指向切片。切片本身已经是一个引用类型，所以它本身就是一个指针！  
1. 切片的声明。格式`var SliceID []TYPE`。MapID为该切片的名字，[]中括号内不指明长度  
2. 切片的初始化。格式`var SliceID []TYPE = ARR[START:END]`。切片通过数组ARR从START号索引到END-1号索引之间的元素构成自己(切分数组，`START:END`被称为slice表达式)。切片的长度为`END-START`，切片的容量为`从所引用的数组索引号START这个元素到这个数组最后一个元素的所有元素的个数`  
   1. e.g.如果定义了一个数组`var arr1 [7]int`  
      1. `var slice1 []type = arr1[2:5]`。`slice1[0]`等于`arr1[2]`。`len(slice1)`切片长度为3,`cap(slice1)`切片容量为5  
      2. `var slice2 []type = arr1[:]`  切片slice2等于完整的arr1数组。另一种表示方式：`slice2 = &arr1`  
      3. `arr1[2:]`和`arr1[2:len(arr1)]`相同，表示包含了数组的2号索引到最后最一个索引的所有元素。  
      4. `arr1[:3]`和`arr1[0:3]`  相同，表示包含了从数组的0号索引到2号索引的所有元素(不包括3号索引的元素)。  
      5. `s := [3]int{1,2,3}[:]`和`s := []int{1,2,3}`相同，表示由数字1、2、3组成的切片  
      6. `s2 := s[:]`使用切片组成的切片，拥有相同的元素，但是仍然指向相同的相关数组  
      7. `var x = []int{2,3,4,5,11}`  创建了一个长度为5的数组且创建了一个相关切片。  
3. 使用`make()`函数来创建一个切片，同时创建好相关数组。  
格式：`var SliceID []TYPE = make([]TYPE, LEN, CAP)`也可以简写为`SliceD ：= make([]TYPE, LEN, CAP)`（其中CAP是可省的，默认与LEN相同）。`MapID`为切片名；`TYPE`为该切片的数据类型；`LEN`为该切片的长度；`CAP`为该切片的总容量。CAP可以理解为切片所引用的数组的长度，切片的长度不能超过容量i.e.不能超过所引用的数组的长度。make()函数接受2个参数：元素的类型、切片的个数。  
   1. e.g.`s2 := make([]int, 10)`。定义了一个名为s2，长度与容量都为10的整型切片  
      1. 这个例子可以拆解为两句，首先会声明一个数组`var XX [10]int`，然后使用该数组初始化一个切片`var s2 []int = XX[:]`  
Note:使用make，而不是直接使用`var persons []Person`的声明方式。还是有所差别的，使用make的方式，当数组切片没有元素的时候，Json会返回[]。如果直接声明，json会返回null  
4. for-range结构  
使用该结构可以对数组或切片中的索引和元素的值进行相关操作，该结构可以返回索引与元素的值  
格式：`for INDEX,VAL := range SliceID {...}`。返回值INDEX为数组或切片的索引；返回值VAL为该索引位置的值；MapID为该数组或切片的名字  
代码示例：[for-range.go](/4.ArraysSlicesMaps/for-range.go)  
5. reslice切片的重组  
在使用make()函数创建切片的时候，LEN作为切片的初始长度，而CAP作为所切片所相关的数组的长度。这么做的好处是切片在达到LEN所定义的上限后，可以扩容，直到扩容到CAP定义的容量。  
改变切片长度的过程称之为**切片的重组(reslicinig)**。  
6. 切片的追加append()与复制copy()  
append格式1：`SLICE2 := append(SLICE1, X1, X2...)`把X1,X2等元素追加给切片SLICE1，追加的元素必须和原切片的元素类型相同。如果SLICE1的容量不足以存储新增的元素，append会分配新的切片来保证已有切片元素和新增元素的存储。因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了。  
append格式2：`SLICE2 := append(SLICE1, SLICE3)`可以将切片SLICE3追加到SLICE1的后面  

copy格式：`copy(SLICE1, SLICE2)`把切片2复制给切片1  

## Maps映射   
map是**键值对**(key-value pairs)的无序集合。这种结构也称**关联数组**(associative array)、**字典**(dictionary)、**散列表/哈希表**(hash table)。这是一种快速寻找值的理想结构：给定key，对应的value可以迅速定位。代码示例：[maps.go](/4.ArraysSlicesMaps/maps.go)  
1. map的声明。格式：`var MapMapID map[KeyType]ValueType`
2. map的初始化。格式：`MapID = make(map[KeyType]ValueType)`。也可以直接在声明时初始化`MapID ：= make(map[KeyType]ValueType)`  
3. map的使用。格式：`MapID[KEY] = VAL`或者`MapID := map[KeyType]ValType{KEY1:VAL1, KEY2:VAL2 ...,KEYn:VALn}`也可以每个键值单独使用一行，但是最后一行要有一个逗号
   1. length长度。格式：`len(MapID)`。map的长度指的是键值对的个数，有几个键值对，长度就是几
   2. key/value键值。格式：`MapID["KEY"]`。引用MapID的某个KEY，可以获取KEY对应的VAL。
5. key/value的删除。格式：`delete(MapID, KEY)`。删除MapID这个map的KEY以及对应的VAL



