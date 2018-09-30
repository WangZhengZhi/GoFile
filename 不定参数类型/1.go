package main

import "fmt"

func main() {
	MyFunction(1, 1, 2, 3, 4, 5, 6)
	test(1, 2,3, 4)
}
func MyFunction(num int, nums ...int) {
	fmt.Printf("%d\n", num)
	for i, data := range nums {
		fmt.Printf("第%d位是 %d\n", i, data)
	}

} //不定参数只能在函数最后一位

func test(args ...int) {
	//全部参数给另外一个函数使用
	//testmyfunc(args...)

	//只想传递后2个参数给另外一个函数使用
	testmyfunc1(args[2:]...)
}
func testmyfunc(temp ...int) {
	for _, data := range temp {
		fmt.Println("data:=", data)
	}
}
func testmyfunc1(temp ...int) {
	for _, data := range temp {
		fmt.Println("data:=", data)
	}
}
