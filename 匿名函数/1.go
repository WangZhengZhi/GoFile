package main

import (
	"fmt"
)

//匿名函数与闭包
func main() {
	a := 10
	str := "hello"
	/*匿名函数方法一*/
	func() {
		fmt.Println(a, str)
	}() //定义匿名函数并使用
	/*方法二*/
	f1 := func() {
		fmt.Println(a, str)
	} //声明匿名函数
	f1() //使用匿名函数
	/*方法三*/
	type FuncType func()
	var f2 FuncType
	f2 = f1 //f2与f1是一样的
	f2()    //调用匿名函数
	/*带有参数的匿名函数的调用:方法一*/
	/*先定义后使用*/
	f3 := func(numa, numb int) {
		fmt.Println(numa, numb)
	}
	f3(1, 2) //调用的时候必须传递参数
	/*带有参数的匿名函数的调用:方法二*/
	/*定义并且使用*/
	func(numa, numb int) {
		fmt.Println(numa, numb)
	}(3, 4)
	/*带有参数以及返回类型的匿名函数*/
	func(numa, numb int) (max int) {
		if numa > numb {
			max = numa
		} else {
			max = numb
		}
		fmt.Println(max)
		return max

	}(10, 20)

}
