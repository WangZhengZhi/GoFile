package main

import "fmt"

/*
结构体指针匿名字段
*/
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	*Person
	addr string
	id   int
}

func main() {
	s := Student{&Person{name: "wang", sex: 'M', age: 18}, "beijing", 6666}
	fmt.Printf("S:%+v\n", s) //打印的指针会是地址
	fmt.Println(s.name, s.sex, s.age)
	var s2 Student
	s2.Person = new(Person) //给指针结构体重新赋值，分配空间
	s2.name = "yoyo"
	s2.sex = 'G'
	s2.age = 18
	s2.addr = "shanghai"
	fmt.Println(s2.name, s2.age, s2.sex, s2.addr)
	var s3 Student
	s3.Person = new(Person) //如果不分配空间会使程序引发错误直接修改的是地址

}
