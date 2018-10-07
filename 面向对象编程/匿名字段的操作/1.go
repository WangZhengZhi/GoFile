package main

import "fmt"

/*
匿名字段的操作
成员的操作
*/
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s := Student{Person{"wang", 'M', 21}, 1, "beijing"}
	//初始化
	/*
		s1:=Student{Person{name:"wang",sex:'M',age:21},1,"shanghai"}

		或者这样的写法
	*/
	s.name = "yoyo"
	s.sex = 'M'
	s.id = 666
	s.age = 18
	s.Person = Person{"go", 'M', 21}
	fmt.Println(s.name, s.sex, s.addr, s.id)
}
