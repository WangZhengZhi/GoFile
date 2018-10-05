package main

import "fmt"

/*
数组指针做函数的参数
 */
func test(a *[5]int) {
	(*a)[0] = 666            //注意格式
	fmt.Println("test:", *a) //打印指向a的内存
}
func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	test(&a)
	fmt.Println("main:", a)
	/*
	最终会改变main 函数内的值，
	 */
}
