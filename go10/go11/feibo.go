package main

import (
	"fmt"
)

func main() {

	for _, v := range feibo(5) {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("%d\n", max(1, 5))
	fmt.Printf("%d\n", min(1, 5)

}
func feibo(value int) []int {
	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x
}

// php python golang node.js
//java /c#
func max(numa int, numb int) (maxnum int) {
	if numa > numb {
		maxnum = numa
	} else {
		maxnum = numb
	}
	return
}
func min(numa int, numb int) (minnum int) {
	if numa > numb {
		minnum = numb
	} else {
		minnum = numa
	}
	return
}
