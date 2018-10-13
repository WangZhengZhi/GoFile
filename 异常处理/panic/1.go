package main

import "fmt"

/*panic问题
遇到不可以恢复的错误状态的时候，数组越界，空指针引用，会引起panic异常，程序崩溃*/
/*人为调用*/
func testa() {
	fmt.Println("a")
}
func testb() {
	fmt.Println("b")
	panic("panic test") /*加入了panic*/
}
func testc() {
	fmt.Println("c")
}
func main() {
	testa()
	testb() /*panic调用导致程序中断*/
	testc()

}
