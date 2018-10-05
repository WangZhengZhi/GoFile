package main

import "fmt"

//闭包特点
//闭包改变变量（内部）会造成变量（内部）改变
func test() int {
	var x int //并没有初始化。值为0
	x++
	return x * x
} //普通函数
/*函数的返回值是一个匿名函数，返回一个函数类型*/
func test01() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	a := 10
	str := "mike"
	func() {
		//闭包使用引用的方式捕获外部变量
		a = 666
		str = "go"
		fmt.Printf("内部：a=%d str=%s\n", a, str)
	}() //()直接调用
	fmt.Printf("外部：a=%d str=%s\n", a, str)
	fmt.Println(test()) //值为1
	fmt.Println(test()) //值为1
	fmt.Println(test()) //值为1
	fmt.Println(test()) //值为1
	fmt.Println(test()) //值为1
	f := test01()
	fmt.Println(f()) //x=1
	fmt.Println(f()) //x=2
	fmt.Println(f()) //x=3
	fmt.Println(f()) //x=4
	fmt.Println(f()) //x=5
	/*
		返回值为一个匿名函数，返回一个函数类型通过f来调用返回的匿名函数
		它不关心这些捕获的常量变量是否已经超出作用域
		只要有闭包使用 变量便会存在
	*/
}
