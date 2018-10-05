package main

import FMT "fmt" //给包起别名
//忽略此包
//import _"fmt"
//为使用init函数
func main() {
	/*
		import ."fmt"
		//调用函数的时候不需要包名字
	*/
	FMT.Printf("hello")
}
