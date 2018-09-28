package main

import (
	"fmt"
)

func main() {
	var a int
	a = 0
	for a < 20 {
		fmt.Printf("%d\n", a)
		a++
		if a > 15 {
			break
		} //break只用用在 switch/ for循环之中
	}

}
