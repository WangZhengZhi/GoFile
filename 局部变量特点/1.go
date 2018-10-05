package main

import "fmt"

func test() {
	a := 10
	fmt.Printf("%d\n", a)
}
func main() {
	//定义在大括号里的变量就是局部变量，只能在{}里有效
	//执行到定义变量那句话，才开始分配空间，离开作用域自动释放
	//作用域，变量起作用的范围
	{
		i := 10
		fmt.Println(i)
	}
}
