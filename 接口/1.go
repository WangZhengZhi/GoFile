package main

import "fmt"

/*
go语言中的接口
定义接口
type interface_name interface{
method_name1 [return type]
.......
}
定义结构体
type struct_name struct{
......}
实现接口的方法
func (struct_name_varible struct_name method_name1()[return_type])
func (struct_name_varible struct_name method_name2()[return_type])
..........




*/

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(Iphone)
	phone.call()
	var man Man
	man = new(Woman)
	fmt.Println(man.id())
	fmt.Println(man.name())
	fmt.Println(man.sex())
}

type Phone interface {
	call()
}
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am nokia ,i call u")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am iphone i call u")
}

type Man interface {
	name() string
	id() int
	sex() string
}
type Woman struct {
}

func (woman Woman) name() string {
	return "zhengzhi_wang"
}
func (woman Woman) id() int {
	return 15426037
}
func (woman Woman) sex() string {
	return "Boy"
}
