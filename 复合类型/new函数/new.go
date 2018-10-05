package main

import "fmt"

//new函数如何使用
func main() {
	var p *int
	//指向一个合法的内存
	p=new(int)
	//申请一个空间内存
	//可以直接操作*p
	*p=666//可以直接赋值操作指针
	fmt.Println("*p=",*p)
	//GO语言不需要释放内存
	q:=new(int)
	*q=777
	fmt.Println("*q=",*q)
	//另外一种形式，可以直接使用，和前面的类似

}
