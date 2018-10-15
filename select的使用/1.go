package main

import "fmt"

/*select的使用*/
/*select实现斐波那契数列*/
func fibonacci(ch chan<- int, flag <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x: //写入
			x, y = y, x+y
		case <-flag: //读取
			fmt.Println("flag is quit:", flag)
			return

		}
	}
}
func main() {
	ch := make(chan int)    //数据通信
	flag := make(chan bool) //判断程序是否结束
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-ch)
		}
		flag <- true
	}()
	fibonacci(ch, flag)
}
