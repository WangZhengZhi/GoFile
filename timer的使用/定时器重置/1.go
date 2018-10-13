package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 10)
	timer.Reset(time.Second * 1) //重新设置为1S
	<-timer.C
	fmt.Println("时间到")
}
