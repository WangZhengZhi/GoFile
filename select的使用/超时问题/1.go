package main

import (
	"fmt"
	"time"
)

/*超时问题,通过select实现*/
func main() {
	ch := make(chan int)
	flag := make(chan bool)
	//新开一个协程
	go func() {
		for {
			select {
			case <-ch: //读取数据
				fmt.Println(<-ch)
			case <-time.After(time.Second * 3):
				fmt.Println("程序超时")
				flag <- true //3秒延迟之后开始写入数据
				break

			}
		}
	}()
	//ch <- 666

	/*for i := 0; i < 5; i++ {
		ch <- i
		//time.Sleep(time.Second * 1)

	} //给ch写数据
	*/
	<-flag
	fmt.Println("程序结束")
}
