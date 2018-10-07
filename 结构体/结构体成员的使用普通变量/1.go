package main

import "fmt"

/*
结构体成员的使用，普通变量
 */

type Student struct {
	id   int
	name string
	sex  byte
	addr string
	age  int
} //结构体声明

func main() {
	//定义一个结构体，普通变量
	var s1 Student
	//操作结构体成员需要用.运算符号
	s1.id = 1
	s1.sex = 'M'
	s1.name = "hello"
	s1.addr = "BeiJing"
	s1.age = 21
	fmt.Println("s1	",s1)
	fmt.Println("s1.name	",s1.name)
}
