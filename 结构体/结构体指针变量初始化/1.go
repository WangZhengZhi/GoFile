package main

import "fmt"

/*
结构体指针变量初始化
 */
type Student struct {
	id   int
	sex  byte
	name string
	age  int
	addr string
} //声明结构体

func main() {
	var p1 *Student = &Student{1, 'M', "tom", 21, "BeiJing"} //顺序初始化
	//别忘了&符号
	fmt.Println("*p1", *p1)
	p2 := &Student{name: "tim", sex: 'M', id: 19} //指定成员初始化
	//别忘了&符号
	fmt.Println("*p2", *p2)
	fmt.Printf("p2 type is %T\n", p2)                            //打印类型
	var p3 *Student = &Student{10, 'M', "hello", 21, "ShangHai"} //顺序初始化
	fmt.Println("*p3 ", *p3)
}
