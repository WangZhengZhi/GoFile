package main

import (
	"fmt"
)

func main() {
	/*
		打印1-100的数字
		如果是3的倍数就打印FIZZ
		如果是 5的倍数 就打印BUZZ
		如果同时是3的倍数也是5的倍数 打印 FIZZBUZZ
	*/
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FIZZBUZZ")
		}
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("FIZZ")
		}
		if i%3 != 0 && i%5 == 0 {
			fmt.Println("BUZZ")
		} else {
			fmt.Println(i)
		} //else必须是在两个}{中间

	}
}
