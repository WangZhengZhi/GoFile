package main

import (
	"fmt"
	"time"
)

//有缓冲的channel

func main() {
	ch := make(chan int, 3)
	fmt.Println("len(ch)=", len(ch)) //缓冲区域剩余的个数
	fmt.Println("cap(ch)=", cap(ch)) //容量
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("i:", i)
			ch <- i //i的值写入channel

		}
		defer fmt.Println("子协程结束")
	}()
	//延时处理
	time.Sleep(time.Second * 1)

	for i := 0; i < 3; i++ {
		num := <-ch //读取chan的值
		fmt.Println("num", num)
	}

}
