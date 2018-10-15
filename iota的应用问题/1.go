package main

import "fmt"

/*iota的应用问题*/
const (
	a = iota //0
	b        //1
	c = "zz" //zz
	d        //zz
	e = iota //4
)

var (
	size     = 1024
	max_size = size * 2
)

func main() {
	fmt.Println(a, b, c, d, e)
	fmt.Println(size, max_size)
	const x = 100
	//fmt.Println(&x,x)//无法得到x的地址

}
