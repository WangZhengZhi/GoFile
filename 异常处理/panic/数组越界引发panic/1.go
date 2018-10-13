package main

import "fmt"

/*数组越界引发panic*/
func testa() {
	fmt.Println("a")
}
func testb(x int) {
	var array [10]int
	array[x] = 111 /*数组索引超出10则会引发panic*/
}
func testc() {
	fmt.Println("c")
}
func main() {
	testa()
	testb(20) /*程序内部会自动调用panic函数*/
	testc()
}
