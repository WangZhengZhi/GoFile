package main

import "fmt"

/*
结构体做函数参数
值传递，引用传递
*/
type Student struct {
	id   int
	name string
	sex  byte
	addr string
}

func test(s Student) {
	s.id = 999
	fmt.Println("test():", s)

}
func main() {
	s := Student{1, "wang", 'M', "DaLian"}
	test(s)
	//形参无法改实参数，函数内部无法改内部实参
	//值传递
	fmt.Println("main:", s)

}
