package main

//函数递归调用的流程

//import库引入时前面加一个.import ."fmt"

//再次调用的时候就不需要写import包的内容

func main() {
	test(3)
	println("main")
}
func test(a int) {
	if a == 1 {
		//函数终止调用的条件
		println("a=", a)
		return //终止函数调用
	}
	test(a - 1)
	println("a=", a) //
}
