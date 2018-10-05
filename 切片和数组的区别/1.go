package main

import "fmt"

/*
切片和数组的区别
 */
func main() {
	a := [5]int{}
	//数组的[]是固定的一个常量，数组不能修改长度，len cap永远是固定的
	fmt.Printf("a::len(a)=%d cap(a)=%d\n", len(a), cap(a))
	//切片，[]里面是空的，或者为...切片的长度/容量可以不固定
	s := []int{}
	fmt.Printf("s::len(s)=%d cap(s)=%d\n", len(s), cap(s))
	s = append(s, 1111) //给切片成员末尾追加一个成员
	fmt.Printf("s append ::len(s)=%d cap(s)=%d\n", len(s), cap(s))

}
