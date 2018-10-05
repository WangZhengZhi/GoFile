package main

import (
	"fmt"
)

//延迟调用
//defer函数
//类似析构函数
/*
defer只能出现在函数或者方法的内部
*/
func test(x int) {
	result := 100 / x
	fmt.Println(result)
}
func main() {
	/*defer fmt.Println("AAAAA")
	//defer会延迟调用
	fmt.Println("BBBB")*/
	/*
		输出为BBBB
		AAAAA
	*/
	/*
		多个defer
		先进后出
		哪怕函数或者某个延迟调用发生错误，调用依旧会被执行
	*/
	defer fmt.Println("第一个defer")
	defer fmt.Println("第二个defer")
	defer fmt.Println("第三个defer")
	defer test(0)
	defer fmt.Println("第四个defer")
}
