package main

import "fmt"

//数组的基本操作
func main() {
	var a [10]int
	//数组元素定义必须是常量
	var b [5]int
	//[数字]，作为数组元素个数
	fmt.Printf("len(a)=%d,len(b)=%d\n", len(a), len(b))
	//注意：数组元素个数必须是常量
	/*n:=10
	var c[n]int
	fmt.Println(c)
	*/
	/*
	数组元素声明的时候，[数字]，数字必须是常量
	 */
	//操作数组元素从0开始到len()-1,不对称原则,这个数字叫下标
	//这时候，下标可以是变量或者是常量
	a[0] = 1
	i := 1
	a[i] = 2 //a[1]=2
	/*
	给数组赋值
	 */
	for i := 0; i < len(b); i++ {
		b[i] = i + 1
	}
	/*
	打印数组
	第一个返回下标，第二个返回元素
	 */
	for key, value := range b {
		fmt.Printf("b[%d]=%d\n", key, value)
	}

}
