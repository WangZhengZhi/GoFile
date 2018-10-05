package main

import "fmt"

//指针作为函数参数
func main() {
	var a int = 10
	var b int = 20
	fmt.Printf("a交换之前的值：%d\n", a)
	fmt.Printf("b交换之前的值：%d\n", b)
	/*
		开始执行交换函数
	*/
	swap(&a, &b)
	fmt.Printf("a交换之后的值：%d\n", a)
	fmt.Printf("b交换之后的值：%d\n", b)
	c := 10
	d := 100
	swap1(&c, &d)
	fmt.Printf("c交换之后的值：%d\n", c)
	fmt.Printf("d交换之后的值：%d\n", d)

}
func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
	return

}
func swap1(x, y *int) { *x, *y = *y, *x } //更加符合go的语言特性
