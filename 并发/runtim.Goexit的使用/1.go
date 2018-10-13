package main

import (
	"fmt"
	"runtime"
)

/*runtime.Goexit的使用*/
func Test() {
	defer fmt.Println("cccccccccc")
	//return //终止此函数
	runtime.Goexit() //终止所在的协程
	fmt.Println("dddddddddddd")
}
func main() {

	go func() {
		fmt.Println("aaaaaaaa")
		//调用函数
		Test()

		fmt.Println("bbbbb")
	}()

	//特地写一个死循环，目的不让主协程结束
	for {
		fmt.Print("")
	}
}
