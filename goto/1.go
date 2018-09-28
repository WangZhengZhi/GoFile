package main

import (
	"fmt"
)

//goto语句的使用

func main() {
	for i := 0; i < 5; i++ {

		if i == 5 {
			goto J
		}
		fmt.Printf("%d\n", i)
	}
J:
	fmt.Println("这是跳转点")
}
