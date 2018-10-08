package main

import "fmt"

/*
接口的练习题目

*/
type Humaner interface {
	SayHi()
} /*定义接口*/
type Person struct {
	name string
	age  int
}

func (p *Person) SayHi() {
	fmt.Println("say hi", p.name, p.age)
}

/*person类实现接口*/
type Student struct {
	Person
	addr string
}

func (p *Student) SayHi() {
	fmt.Println("say hi", p.age, p.name, p.addr)
} /*student类实现接口*/

type Mystr string

func (tmp Mystr) SayHi() {
	fmt.Println(tmp)
} /*mystr（自定义类型）实现接口*/

func main() {
	p := &Person{"wang", 19}
	var i Humaner
	i = p
	i.SayHi()
	var str Mystr = "hello mike"
	i = &str
	i.SayHi()
	s := &Student{Person{"li", 20}, "bj"}
	i = s
	i.SayHi()
}
