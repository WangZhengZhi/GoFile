package main

import (
	"fmt"
	"time"
)

/*通过channel实现数据交互*/
func main() {
	ch := make(chan string)
	go func() {
		defer fmt.Println("我是子协程调用完毕")
		for i := 0; i < 3; i++ {
			fmt.Println("i=", i)
			time.Sleep(time.Second * 1)

		}
		ch <- "我是子协程，工作完毕"
	}()
	str := <-ch
	fmt.Println(str)
	defer fmt.Println("我是主协程,工作完毕")

}
