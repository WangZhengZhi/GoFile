package main

import . "fmt"

//多态以及回调函数
//回调函数，函数参数有一个是函数类型,就是回调函数
//
type FuncType func(int, int) int

//计算器可以进行四则运算
//多态，多种形态.调用同一个接口，不同的表现，可以实现不同表现加减乘除
//先有有想法，后面再实现功能
//实现加法计算
func Add(numa, numb int) int {
	return numb + numa
}
func Minus(numa, numb int) int {
	return numa - numb

}
func Calc(a, b int, Ftest FuncType) (result int) {
	Println("Calc")
	result = Ftest(a, b) //这个函数并没有实现
	//result=Add(a,b)必须是先定义才能调用,函数不能扩展,函数写死了
	return

}
func main() {
	a := Calc(1, 2, Add)
	println("a=", a)
	b := Calc(1, 1, Minus)
	println(b)

}
