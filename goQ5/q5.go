package main

import (
	"fmt"
)

func main() {

	addfloat64()
	swapnum(2, 7)
	swapnum(7, 2)
	pop(10)

}
func addfloat64() {
	sum := 0.0
	a := []float64{1, 2, 3, 4}
	xs := a[:]
	var avg float64
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))

	}
	fmt.Printf("平均值是%f\n", avg)
} //函数计算float64平均值Q5
func swapnum(numa int, numb int) {
	fmt.Printf("交换之前是:%d,%d\n", numa, numb)
	fmt.Printf("交换之后是:%d,%d\n", numb, numa)
} //函数交换数据Q6
func pop(index int) {
	var list [1000]int
	var i int
	for i = 0; i < len(list) && i < index; i++ {
		list[i] = i
		fmt.Printf("list[%d]=%d\n", i, list[i])
	}

} //入栈函数
func push() {
	//出栈函数
}
