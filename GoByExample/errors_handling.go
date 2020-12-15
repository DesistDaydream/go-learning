// Go语言使用一个独立、明确的返回值来传递错误信息。
package main

import (
	"errors"
	"fmt"
)

// 错误通常是最后一个返回值，并且是`error`类型，一个内建的接口
func f1(arg int) (int, error) {
	if arg == 42 {
		// `errors.New`构造一个给定的错误信息的基本`error`值
		return -1, errors.New("不能使用42工作")
	}
	// 返回的错误值为`nil`，则表示没有错误
	return arg + 3, nil
}

// 通过实现`Error`方法来自定义`error`类型，这里使用自定义错误类型来表示
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func errorHandling() {
	// 循环测试各个返回错误的函数
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	// 如果想在程序中使用一个自定义错误类型中的数据
	// 需要通过`类型断言`来得到这个错误类型的实例
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
