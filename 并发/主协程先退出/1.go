package main

import (
	"fmt"
	"time"
)

/*主协程先退出

goroutine先退出
主协程退出，其他协程也要退出

*/
func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程  i:=", i)
			time.Sleep(time.Second * 1)
			if i == 2 {
				break
			}
		}

	}()
	i := 0
	for {
		i++
		fmt.Println("main  i:=", i)
		time.Sleep(time.Second * 1)
		if i == 2 {
			break
		}
	}
}
