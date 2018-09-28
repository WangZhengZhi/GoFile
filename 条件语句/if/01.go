package main

import (
	"fmt"
)

func main() {
	var num int
	num = 100
	if num >= 100 {
		fmt.Println(">=100")
	} else {
		fmt.Println("<100")
	}
	var numa int
	var numb int
	numa = 100
	numb = 200
	if numa == 100 {
		if numb == 200 {
			fmt.Printf("numa=%d,numb%d\n", numa, numb)
		}
	}
	fmt.Printf("条件控制语句结束\n")

}
