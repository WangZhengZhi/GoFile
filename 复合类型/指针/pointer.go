package main

import "fmt"

func main() {
	var a int = 10
	//每个变量有2层含义：变量的内存，变量的地址
	fmt.Printf("a=%d\n", a)   //变量的内存
	fmt.Printf("&a=%v\n", &a) //变量的地址
	var p *int                //定义一个变量P 为*int
	p = &a                    //将a的地址赋值给指针P
	fmt.Printf("p=%v,&a=%v\n", p, &a)
	*p = 666
	fmt.Printf("*p=%v,a=%d\n", *p, a)
}
