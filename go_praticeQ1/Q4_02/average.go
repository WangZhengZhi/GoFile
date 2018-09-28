package main

import (
	"fmt"
)

//计算float64类型的平均值

func main() {
	var avg float64
	sum := 0.0
	str := [...]float64{1, 2, 3, 4, 5, 6, 7}
	slice := str[:]
	switch len(slice) {
	case 0:
		avg = 0
	default:
		for _, v := range slice {
			sum += v
		}
		avg = sum / float64(len(slice))

	}
	fmt.Println(avg)

}
