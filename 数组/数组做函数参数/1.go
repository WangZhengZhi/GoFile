package main

import "fmt"

/*
数组做函数参数，是值传递
实参数组的每个元素给形参数组COPY一份
 */
func test(a [5]int) {
	a[0] = 666
	fmt.Println("test:", a)

}
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	test(a) //数组传递
	//数组传递会影响效率，因为是COPY过程
	//数组越大越浪费时间
	//形参数组是实参数组的赋值品
	fmt.Println("main:", a) //不会对数组产生影响
}
