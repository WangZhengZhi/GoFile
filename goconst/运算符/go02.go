package main

import (
	"fmt"
)

func main() {
	var A bool
	var B bool
	A = true
	B = false
	fmt.Println(A && B)
	fmt.Println(A || B)
	fmt.Println((!(A && B)))
	//位运算符与 或 异或
	//& | ^
	//左移 右移
	//<< >>
	var a uint
	var b uint

	a = 0
	b = 1
	var c uint
	c = 7
	fmt.Println(a & b)  //0
	fmt.Println(a | b)  //1
	fmt.Println(a ^ b)  //1
	fmt.Println(b << 1) //1 左移动 1位10=2
	fmt.Println(c >> 1) //111>>1 011=3
	fmt.Println(c << 1) //111<<1 1110=14
	/*
		指针运算符
		&	返回变量的存储地址，
		&a	将给出变量的实际地址
		*   指针变量
		*a	是一个指针变量
	*/
	var numa int = 4
	var numb int32
	var numc float64

	fmt.Printf("numa的类型%T\n", numa)
	fmt.Printf("numa的类型%T\n", numb)
	fmt.Printf("numa的类型%T\n", numc)
	var ptr *int
	ptr = &numa
	fmt.Printf("a的值%d\n", numa)
	fmt.Printf("*ptr的值：%d", *ptr) //*ptr的值
	fmt.Printf("*ptr 的内存地址%d\n", ptr)
	/*
		运算符的优先级别
		与数学是一致的
	*/
	var numx int
	var numy int
	var sum int
	numx = 10
	numy = 10
	sum = 0
	sum = numx + numy*10
	fmt.Printf("%d\n", sum)
	sum=(numx+numy)*(numx*numy)//(10+10)*(10*10)=20*100
	fmt.Printf("%d\n",sum)

}
