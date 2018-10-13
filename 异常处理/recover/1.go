package main

import "fmt"

/*recover的使用
要配合defer函数*/
func testa() {
	fmt.Println("a")
}
func testb(x int) {
	/*设置defer*/
	defer func() {
		//recover()
		//fmt.Println("错误信息", recover()) /*可以打印panic信息*/
		if err := recover(); err != nil {
			fmt.Println(err)
		} /*如果产生异常则打印错误信息*/
	}() /*别忘了（），调用匿名函数*/
	var array [10]int
	array[x] = 1111
}
func testc() {
	fmt.Println("c")
}
func main() {
	testa()
	testb(20)
	testc()
}
