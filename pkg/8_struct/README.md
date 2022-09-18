# 概述

# Maps 映射

map 是**键值对**(key-value pairs)的无序集合。这种结构也称**关联数组**(associative array)、**字典**(dictionary)、**散列表/哈希表**(hash table)。这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。代码示例：[maps.go](/4.arrays_slices_maps/maps.go)

1. map 的声明。格式：`var MapMapID map[KeyType]ValueType`
2. map 的初始化。格式：`MapID = make(map[KeyType]ValueType)`。也可以直接在声明时初始化`MapID ：= make(map[KeyType]ValueType)`
3. map 的使用。格式：`MapID[KEY] = VAL`或者`MapID := map[KeyType]ValType{KEY1:VAL1, KEY2:VAL2 ...,KEYn:VALn}`也可以每个键值单独使用一行，但是最后一行要有一个逗号
   1. length 长度。格式：`len(MapID)`。map 的长度指的是键值对的个数，有几个键值对，长度就是几
   2. key/value 键值。格式：`MapID["KEY"]`。引用 MapID 的某个 KEY，可以获取 KEY 对应的 VAL。
4. key/value 的删除。格式：`delete(MapID, KEY)`。删除 MapID 这个 map 的 KEY 以及对应的 VAL

# Structs 结构体

**Structs(结构体)** 是一种复合数据类型，当需要一个[自定义 Type](/1.Type/README.md)，且这个 Type 由一系列的属性组成，每个属性都有自己的 **类型** 和 **值** 的时候，就需要使用 Structs，Structs 把数据聚集在一起，然后访问这些数据的时候，好像这些数据是一个独立实体的一部分。Structs 也是值类型，可以通过 `new()` 函数创建。

组成结构体的属性分两部分

- FIELD(字段)
- BaseType(基础类型)
  > 基础类型可以是另一个结构体,表示该结构体包含另一个结构体

每个字段都有其对应的基础数据类型，在一个结构体中，FIELD 名字必须是唯一的。代码示例：[struct.go](/7.structs_and_interfaces/struct/struct.go)
结构体的定义格式：

```go
type StructID struct {
	FIELD1 BaseType1 ["TAG"]
	FIELD2 BaseType2 [`TAG`]
	...
}
```

也可以使用简单的方法定义：`type T struct {a, b int}`

结构体的初始化格式：`VarID ：= new(PackageID.StructID)`。若初始化的结构体为当前包的，则可以省略 PackageID

结构体中字段的引用格式：`StrcutID.FIELD1`。结构体名，中间跟一个点，再接该结构体内的字段名。
在 Go 语言中，这个`.`点符号叫做**选择器(selector)**。无论定义的变量是一个结构体类型还是一个结构体类型指针，都是用同样的**选择器符(selector-notation)**来引用结构体的字段。

TAG
除了 FIeld 和 BaseType 之外，还可以给该属性添加 TAG(标签)，TAG 使用`双引号`或者`重音号`来表示。这些 TAG 能被用来做文档或者重要的标签。
TAG 里面的内容在正常编程中没有作用。一般在反射、某些第三方库(比如 gin 的数据绑定功能)、等等地方可以起到关键的作用。
