package main

//go语言使用值传递
import "fmt"

func Swap(a, b int) {
	a, b = b, a
	fmt.Printf("Swap: a=%d b=%d\n", a, b)
}
func main() {
	a, b := 10, 20
	Swap(a, b)                           //函数内部交换
	fmt.Printf("main:a=%d b=%d\n", a, b) //main函数并没有交换

}
