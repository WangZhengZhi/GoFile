package main

import (
	"fmt"
)

func main() {

	myfunc(1, 2, 3, 4, 5)
	a := func() {
		println("hello world!!!")
	} //函数作为一个值传递
	a()                   //调用函数
	fmt.Printf("%T\n", a) //打印输出a的类型
	//使用%T打印并查看某个类型
}
func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("the number is :%d:\n", n)
	}
} //arg ...int代表是传入不定数量的参数，参数的类型全是int
func printit(x int) {
	fmt.Printf("%v\n", x)
} //仅打印无返回值
func callback(y int, f func(int)) {
	f(y)
} //回调函数
func throwpanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f()
	return
} //panic实例
