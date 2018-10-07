package main

import "fmt"

/*
匿名字段就是OOP的继承，实现代码复用
匿名字段的初始化
*/
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person //只有类型没有名字，匿名字段，继承了Person的成员
	id     int
	addr   string
}

func main() {
	//顺序初始化
	var s Student = Student{Person{"wang", 'M', 21}, 1, "beiJing"}
	fmt.Println("s:	", s)
	//自动推导类型
	s1 := Student{Person{"wang", 'G', 20}, 2, "shangHai"}
	fmt.Println("s1:	", s1)
	fmt.Printf("s2=%+v\n", s1) //%+v可以使显示更加详细
	//指定成员初始化，没有初始化的成员自动赋值为0
	s3 := Student{Person: Person{name: "zhao"}, id: 1}
	fmt.Printf("s3:%+v\n", s3)
}
