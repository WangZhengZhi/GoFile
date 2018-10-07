package main

import "fmt"

/*
非结构体匿名字段
*/
type mystr string // 自定义类型给一个类型改名字
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person // 结构体匿名字段
	addr   string
	int    //基础类型的匿名字段
	mystr
}

func main() {
	s := Student{Person{name: "wang", sex: 'M', age: 18}, "BJ", 111, "hhhhh"}
	fmt.Printf("%+v\n", s)
	fmt.Println(s.int)
	fmt.Println(s.Person)
	fmt.Println(s.int)
	s.Person = Person{name: "hello", sex: 'M', age: 10}
	fmt.Println(s.Person)
}
