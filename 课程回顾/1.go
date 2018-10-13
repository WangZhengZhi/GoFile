package main

import "fmt"

/*课程回顾*/
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person /*匿名字段*/
	id     int
}

func main() {
	s := Student{}
	s.name = "hello"
	fmt.Println(s)
	var p *Student
	p = &Student{}
	fmt.Println(p) /*指针的操作*/
	p1 := new(Student)
	p1.name = "wang" /*只有.操作*/
	fmt.Println(p1)
	/*同名字段
	就近原则，先找当前作用域
	若要指定s.person.name*/
	/*
			type mystr string
			type person struct{
			mtstr
			}
			s.mystr
		自定义类型
	*/

	/*指针类型
	type Person struct{
	}
	type Student struct{
	*Person
	}
	s:=student{&Person()}

	*/

}
