// 延缓
// defer用来确保一个函数调用在程序执行结束前执行。同样用来执行一些清理工作
package main
import "fmt"

func main() {
    fmt.Println("使用defer效果")
    function10()
    fmt.Println("\n未使用defer效果")
    function11()
    fmt.Println("\n其余defer实例")    
    f()
}

func function10() {
    fmt.Printf("In function1 at the top\n")
    //带有defer关键字的函数调用会推迟到函数返回之前一刻才执行
    defer function2()
    fmt.Printf("In function1 at the bottom!\n")
}

func function11() {
    fmt.Printf("In function1 at the top\n")
    function2()
    fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
    fmt.Printf("function2: Deferred until the end of the calling function!\n")
}

//当有多个defer行为时，会以逆序的顺序执行(类似栈i.e.后进先出)
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

///defer语句实现代码