package main

import (
	"fmt"
	"math"
)

//语言函数作为值
func main() {

	var x float64
	x = 9
	xroot := getxroot(x)
	fmt.Printf("%f root is %f\n", x, xroot)
}
func getxroot(x float64) float64 {
	return math.Sqrt(x)

} //开根号的函数
