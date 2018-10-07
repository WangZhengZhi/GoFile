package main

import "fmt"

/*
面向过程和面向函数的区别
*/
//实现2数相加
//面向过程
func test(a, b int) int {
	return a + b
}

//面向对象，方法：给某一个类型绑定一个函数
type some int

func (temp some) Add(other some) some {
	return temp + other
}
func main() {
	result := test(1, 1)
	fmt.Println(result)
	var a some = 2
	r := a.Add(3)
	fmt.Println(r)
}
