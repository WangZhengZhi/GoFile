package main

import (
	"fmt"
	"unsafe"
)

var (
	e int
	f int
) //全局变量可以声明但不使用

func main() {
	const high = 10
	//const可以声明不使用

	const long = 10
	const width = 10
	a, b, c := 1, 2, "this is string"
	var area int
	area = high * long * width
	fmt.Printf("area=%d\n", area)
	fmt.Println(a, b, c)
	const (
		unknow  = 10
		unknowa = 100
		unknowb = 1000
	) //const 可以用来枚举
	const (
		str1 = "hello"
		str2 = len(str1)
		str3 = unsafe.Sizeof(str1)
	) //len()/unsafe.Sizeof()函数计算表达式的值
	fmt.Println(str1, str2, str3)
	const (
		numa = iota
		numb
		numc
	) //iota是一个特殊常量not one iota出自新约中的短语
	fmt.Println(numa, numb, numc)
	const (
		numd = iota
		nume
		numf = "hahaha"
		numi
		numj = iota //恢复计数和iota行数有关联numj=4
	)
	fmt.Println(numd, nume)
	fmt.Println(numf, numi)
	fmt.Println(numj) //numj=4
	/*iota实验*/
	const (
		d = 1 << iota //转为2进制1=01 左移动零位 还是1
		e = 3 << iota //转化为2进制3=11 左移动一位为110=6
		f = 4 << iota //4=100 左移二位 10000=16
		g             //在4的基础上左移三位4=100, 100000=32
		h             //在4的基础上左移四位 4=100，1000000=64
	)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

}
