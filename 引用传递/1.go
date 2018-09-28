package main

import (
	"fmt"
)

//引用传递
//引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数
func main() {
	a := 10
	b := 20
	swapnum(&a, &b) //
	//使用&字符表示传递的是地址
	//地址交换并不是值的交换b不能直接赋值
	//值传递影响实际值
	fmt.Printf("a=%d,b=%d\n", a, b)
	c, d := 10, 20
	swapnum1(c, d)
	fmt.Printf("c=%d,d=%d\n", c, d)

}
func swapnum(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

} //(a *int,b *int)->(a,b *int)
func swapnum1(a, b int) int {
	var temp int
	temp = a
	a = b
	b = temp
	return temp

} //值传递
