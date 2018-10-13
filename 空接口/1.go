package main

import "fmt"

/*空接口不包含任何的方法
所有类型都实现了空接口
就是一个万能类型，可以保存任何类型的值*/
func test(arg ...interface{}) {
	//可以接受任意类型的参数，任意个数
}
func main() {
	//空接口就是万能类型，可以保存任意类型的值
	var i interface{} = 1
	fmt.Println("i=", i)
	i = "abc"
	fmt.Println("i=", i)
}
