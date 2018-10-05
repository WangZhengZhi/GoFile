package calc

import "fmt"

func init() {
	fmt.Println("这里是被调用函数的init函数")
}
func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}
