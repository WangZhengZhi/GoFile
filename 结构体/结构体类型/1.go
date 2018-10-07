package main

import "fmt"

/*
结构体类型以及普通变量初始化
 */
type Student struct {
	id   int
	name string
	sex  byte //字符类型
	age  int
	addr string
}

func main() {
	var s1 Student = Student{1, "tim", 'M', 19, "Beijing"} //顺序初始化，每个成员必须初始化
	fmt.Println("s1",s1)
	//指定成员初始化，没有初始化的成员自动赋值为0
	s2:=Student{name:"tom",sex:'M'}
	fmt.Println("s2",s2)
}
