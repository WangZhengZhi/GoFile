/*
Go 语言支持匿名函数，可作为闭包。
匿名函数是一个"内联"语句或表达式
匿名函数的优越性在于可以直接使用函数内的变量，不必申明
*/
package main

import (
	"fmt"
)

func main() {
	nextnum := getnumaddone()
	fmt.Printf("%d\n", nextnum())
	fmt.Printf("%d\n", nextnum())
	fmt.Printf("%d\n", nextnum())
	fmt.Printf("%d\n", nextnum())
	nextnum1 := getnumaddone()
	fmt.Printf("%d\n", nextnum1())
	fmt.Printf("%d\n", nextnum1())
	fmt.Printf("%d\n", nextnum1())
	add_func := add(1, 2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())

}
func getnumaddone() func() int {
	i := 0
	return func() int {
		i += 1
		return i

	} //闭包函数，直接使用函数里的变量不用声明
}
func add(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}
} //带参数的闭包函数
