package main

import (
	"fmt"
)

func main() {
	/*
	   01:创建一个for循环循环10次 打印输出的值
	*/
	for i := 0; i < 10; i++ {
		fmt.Println("________", i)
	}
	/*
	   02:用goto 改写01的方式 不用for
	*/
	num02 := 0
J:
	fmt.Println(num02)
	num02++
	if num02 < 10 {
		goto J

	}
	/*
		写一个array
		并且打印array

	*/
	//var array [10]int
	var array [10]int
	for i:=0;i<len(array);i++{
		fmt.Println(array[i])
	}

}
