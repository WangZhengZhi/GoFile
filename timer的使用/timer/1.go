package main

import (
	"fmt"
	"time"
)

//timer的使用
/*验证time.newtimer()，时间到了，只会响应一次*/
func test() {
	timer := time.NewTimer(time.Second * 1)
	for {
		<-timer.C
		fmt.Println("时间到了")
	}
}
func main() {
	//创建一个定时器，设置时间为2s，2s之后往timer通道写内容（当前时间）
	timer := time.NewTimer(time.Second * 2)
	fmt.Println("当前时间", time.Now())
	//2s后写入数据往timer.C写数据
	t := <-timer.C //channel没有数据前会阻塞
	fmt.Println("t:", t)
	//test()
	/*timer.C 只会执行一次，不会一直执行不是循环*/
}
