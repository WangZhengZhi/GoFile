package main

import "fmt"

const MAX int = 3

func main() {
	var a int = 20
	var ip *int
	ip = &a //指针变量的存储地址
	fmt.Printf("a的变量地址是：%x\n", &a)
	fmt.Printf("指针变量的存储地址是%x\n", ip) //指针变量的存储地址
	fmt.Printf("*ip变量的值是：%d\n", *ip) //使用指针访问值
	/*
		go语言中的空指针
	*/
	var ptr *int
	fmt.Printf("prt=%v\n", ptr)
	//判断是不是空指针
	//if(ptr!=nil)
	//if(ptr==nil)
	/*
	 go指针数组的讲解
	*/
	var ptr1 [MAX]*int
	var i int
	arraya := [MAX]int{10, 100, 1000}
	for i = 0; i < MAX; i++ {
		ptr1[i] = &arraya[i]
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d]=%d\n", i, *ptr1[i])
	}

}
