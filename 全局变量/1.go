package main

import "fmt"

//定义在函数外部的变量是全局变量
var a int

func test() {
	fmt.Println(a)
} //全局变量在任何地方都可以使用
func main() {
	a = 10
	fmt.Println(a)
	test()
}
