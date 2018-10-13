package main

import (
	"fmt"
	"time"
)

/*停止定时器*/
func main() {
	timer := time.NewTimer(time.Second * 3)
	go func() {
		<-timer.C
		fmt.Println("子协程可以打印了，定时器时间到了")
	}()
	timer.Stop() //停止定时器

	for {
		fmt.Printf("")
	}
}
