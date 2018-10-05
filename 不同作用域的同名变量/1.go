package main

import "fmt"

var a byte

func main() {
	var a int //局部变量
	//不同作用域允许同名变量
	//使用变量的原则是就近原则
	fmt.Printf("%T\n", a) //int
	{
		var a float64
		fmt.Printf("%T\n", a) //float64

	}
	test()
}
func test() {
	fmt.Printf("%T\n", a) //uint
}
