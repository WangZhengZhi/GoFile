package main

import (
	"fmt"
)

func main() {
	var a int
	for a = 0; a < 10; a++ {
		fmt.Println(a)
	}
	nums := [...]int{1, 2, 3, 4, 5}
	for key, value := range nums {
		fmt.Printf("第%d位是：%d\n", key, value)
	}
	for _, value := range nums {
		if value > 3 {
			fmt.Println(value)
		}
	}
	var x, y int
	for x = 2; x < 20; x++ {
		for y = 2; y <= (x / y); y++ {
			if x%y == 0 {
				break
			}
		}
		if y > (x / y) {
			fmt.Printf("%d:是素数\n", x)
		}
	}
}
