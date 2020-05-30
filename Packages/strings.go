package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(
    //判断test字符串是否包含es，结果为"true"
    strings.Contains("test", "es"),"\n",

    //统计字符串出现的次数，统计t在test中出现的次数，结果为"2"
    strings.Count("test", "t"),"\n",

    //判断test字符串是否以te开头，结果为"true"
    strings.HasPrefix("test", "te"),"\n",

    //判断test字符串是否以st结尾，结果为"true"
    strings.HasSuffix("test", "st"),"\n",

    //定位字符串t在字符串test中第一次出现的位置，结果为0,如果t为多个字符，则第一次的位置按照总字符串中每一个字符的位置算，比如testes中es位置为1而不是0，最后出现位置为4
    //t的位置为0，e的位置为1，s的位置为3，t的位置为"4"
    //若结果为-1则不在test字符串中
    strings.Index("test", "t"),"\n",

    //返回字符串t在字符串test中最后一次出现的位置，结果为"3"
    strings.LastIndex("test", "t"),"\n",

    //拼接slice(切片)到字符串，"a-b"
    strings.Join([]string{"a","b"}, "-"),"\n",

    //重复字符串，重复字符串"a"5次，结果为"aaaaa"
    strings.Repeat("a", 5),"\n",

    //字符串替换，将aaaa字符串中的a替换成b，只替换前两个，结果为"bbaa"
    strings.Replace("aaaa", "a", "b", 2),"\n",

    // []string{"a","b","c","d","e"}
    strings.Split("a-b-c-d-e", "-"),"\n",

    //修改字符串大小写，将指定的字符串中的大写全部变成小写，"test"
    strings.ToLower("TeST"),"\n",

    //修改字符串大小写，将指定的字符串中的大写全部变成大写，"TEST"
    strings.ToUpper("Test"),"\n",
  )
}
