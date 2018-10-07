package main

import "fmt"

/*
结构体做函数参数，引用传递(地址传递)
*/
type Student struct {
	id   int
	name string
	sex  byte
	age  int
}

func test(p *Student) {
	p.id = 999 //操作指针所指向的内存
	fmt.Println("test:", *p)

}

func main() {
	s := Student{1, "wang", 'M', 21}
	test(&s) //地址传递，（引用传递），形参可以改实参 &符号取地址
	fmt.Println("main:", s)

}
