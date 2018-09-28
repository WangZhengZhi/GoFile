package main

import (
	"fmt"
)

//continue语句结束本次语句执行下一句语句

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue //j结束了i==5的语句
		}
		fmt.Printf("%d\n", i)
	}
}
