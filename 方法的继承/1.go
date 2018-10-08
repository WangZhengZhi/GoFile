package main

import "fmt"

/*
方法的继承
*/
type Person struct {
	name string
	sex  byte
	age  int
}

//person 类型实现了一个方法
func (tmp *Person) PrintInfo() {
	fmt.Printf("name=%s sex=%c age=%d", tmp.name, tmp.sex, tmp.age)
}

//有个学生继承了person 成员和方法都继承了
type Student struct {
	Person
	id   int
	addr string
}

func main() {
	s := Student{Person{"mike", 'M', 21}, 666, "BeiJing"}
	s.PrintInfo() //方法的继承

}
