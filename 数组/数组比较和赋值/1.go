package main

import "fmt"

/*
数组的比较与赋值
 */
func main() {
	/*
	数组支持比较；支持== !=
	 */
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2}
	fmt.Println("a==b", a == b)
	fmt.Println("a==c", a == c)
	fmt.Println("a!=c", a != c)
	//相同类型的数组可以赋值
	var d [5]int
	d = a
	fmt.Println("d:",d)


}
