package main

/*
int add(int a, int b) {
	return a + b;
}
*/
import "C"
import "fmt"

func main() {
	result := C.add(3, 4)
	fmt.Println(result)
}
