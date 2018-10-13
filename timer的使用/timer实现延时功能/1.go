package main

import (
	"fmt"
	"time"
)

/*timer实现延时功能,就是sleep（）*/
func main() {
	//延时2S后打印一句话

	//通过time.sleep()实现
	/*fmt.Println("时间开始")
	time.Sleep(time.Second * 1)
	fmt.Println("时间到了")*/
	//通过timer实现
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	fmt.Println("时间到了")
	//定时2s阻塞2s   2S后产生事件   写入内容到channel
	<-time.After(time.Second * 2)

}
