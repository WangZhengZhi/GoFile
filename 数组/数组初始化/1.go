package main

import "fmt"

/*
数组的初始化
 */
func main() {
	/*
	声明的同时赋值，叫做初始化
	 */
	var a [5]int = [5]int{1, 2, 3, 4, 5} //全部初始化
	fmt.Println(a)
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	/*
	部分初始化,没有初始化的元素自动赋值为0
	 */
	c := [5]int{1, 2}
	fmt.Println(c) //输出为 1 2 0 0 0
	/*
	指定某个元素初始初始化
	 */
	d := [5]int{2: 10, 4: 20}
	fmt.Println("d[2]=", d[2])
	fmt.Println("d[4]=", d[4])
}
