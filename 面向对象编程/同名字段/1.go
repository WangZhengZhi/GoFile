package main

import "fmt"

/*
同名字段
*/
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person
	name string //和person同名
	addr string
}

//Student中有和Person一样的同名字段
func main() {
	var s Student
	s.name = "mike" //操作是student，默认操作student
	//如果能在本作用域找到此成员，就操作此成员，如果没有找到就找继承的字段
	s.sex = 18
	s.addr = "BJ"
	fmt.Printf("%+v\n", s)
	//显示调用
	s.Person = Person{name: "wang"}
	fmt.Printf("%+v\n", s)

}
