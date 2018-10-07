package main

import (
	"fmt"
)

/*
结构体比较和赋值
==或者！=
*/
type Student struct {
	id   int
	name string
	sex  byte
}

func main() {
	s := Student{1, "wang", 'M'}
	s1 := Student{1, "wang", 'M'}
	s2 := Student{2, "tom", 'G'}
	fmt.Println("s==s1", s == s1)
	fmt.Println("s1!=s2", s1 != s2)
	/*
		同类型的2个结构体变量可以是互相赋值
	*/
	var temp Student
	temp = s2 //结构体赋值
	fmt.Println("temp", temp)

}
