package main

import "fmt"

/*
指针变量的方法集
*/
type Person struct {
	name string
	age  int
	sex  byte
}

func (p Person) SetValueInfo( /*name string, age int, sex byte*/ ) {
	/*p.name = name
	p.age = age
	p.sex = sex*/
	fmt.Println("set value info") //值传递
}
func (p *Person) SetPointerInfo( /*name string, age int, sex byte*/ ) {
	/*p.age = age
	p.name = name
	p.sex = sex*/
	fmt.Println("set pointer info") //引用传递
}

func main() {
	/*结构体变量是一个指针变量，它能够调用那些方法，这些方法是一个集合，简称方法集*/
	p := &Person{"mike", 18, 'M'}
	p.SetPointerInfo()    //调用set pointer info
	(*p).SetPointerInfo() /*把*p转化为p再调用，等价与上面*/
	p.SetValueInfo()
	//调用set value info
	/*内部做了转化，先把指针P转成*p后再调用*/
	(*p).SetValueInfo() //直接的方式

}
