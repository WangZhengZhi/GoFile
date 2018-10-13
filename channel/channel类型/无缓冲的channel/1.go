package main

import (
	"fmt"
	"time"
)

/*无缓冲的channel 类型
 */
func main() {
	//无缓冲的channel
	ch := make(chan int)
	fmt.Println("len(ch),cap(ch)", len(ch), cap(ch)) //len() 缓冲区剩余数据的个数，cap()缓冲区的大小
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程:", i)
			ch <- i //写入数据
		}
	}()
	//延时
	time.Sleep(time.Second * 2)
	for i := 0; i < 3; i++ {
		num := <-ch //读通道中的内容，没有内容前会阻塞
		fmt.Println("num=", num)

	}

}
