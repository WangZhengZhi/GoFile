package main

import "fmt"

/*关于new的问题*/
func main() {
	//list := new([]int) //错误的写法
	list := make([]int, 0) //make要指定len（）

	list = append(list, 1)
	fmt.Println(list)
	/*s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)*/

	/*
		slice使用切片别忘了'....'
		而且这个是array

	*/

}
