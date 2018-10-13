package main

import (
	"fmt"
	"runtime"
)

/*
runtime.GOMAXPROCESS的使用，设置CPU核数
cpu核数越大，时间片轮转越快
*/
func main() {
	coreNum := runtime.GOMAXPROCS(1) /*指定核数运算*/
	fmt.Println("core num:", coreNum)

	/*for {
		go fmt.Print(1)
		fmt.Print(0)
	}*/
}
