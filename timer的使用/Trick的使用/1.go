package main

import (
	"fmt"
	"time"
)

/*Tricker的使用*/
//Tricker是循环的

func main() {
	Tricker := time.NewTicker(time.Second * 1)
	i := 0
	for {
		<-Tricker.C
		i++
		fmt.Println(i)
		if i == 5 {
			Tricker.Stop() //停止循环
			break
		}
	}
}
