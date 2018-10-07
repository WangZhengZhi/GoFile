package main

import "fmt"

/*
结构体成员的使用，指针变量
 */
type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
} //声明定义一个结构体

func main() {
	/*
	1/指针有合法指向后，才能操作成员

	 */
	var s Student  //先定义一个结构体普通变量
	var p *Student //再定义一个指针变量保存s的地址
	p = &s         //将s的地址给p。不要忘了&
	p.id = 1
	p.name = "hello"
	(*p).addr = "BJ" //和p.addr是一样的，通常写作p.addr
	fmt.Println("*p	",*p)
	/*
	2/通过new申请一个结构体
	 */
	p1 := new(Student)
	p1.addr = "BJ"
	p1.name = "tom"
	p1.age = 19
	p1.sex = 'M'
	p1.id=2
	fmt.Println("p1	",p1)

}
