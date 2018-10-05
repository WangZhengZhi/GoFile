package main

import "fmt"

func main01() {
	a := 10
	b := 20
	defer func() {
		fmt.Println("内部：", a, b)
	}() //延迟调用匿名函数，最后执行
	a = 111
	b = 222
	fmt.Println("外部", a, b)
	/*
		外部 111 222
		内部： 111 222

	*/

}
func main() {
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Println("内部：", a, b)
	}(a, b) //先把参数传递了，只是没有调用
	a = 111
	b = 222
	fmt.Println("外部:", a, b)
	/*

		外部: 111 222
		内部： 10 20
	*/
}
