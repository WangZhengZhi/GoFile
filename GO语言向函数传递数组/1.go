package main

import (
	"fmt"
)

func main() {
	array := []float64{1, 2, 3, 4, 5, 6}
	avg := getarrayavg(array, 6) //数组作为参数传递给函数
	fmt.Printf("avg=%v\n", avg)
	array1 := []int{1, 2, 3, 4, 5, 6}
	avg1 := getintavg(array1, 6) //数组作为参数传递给函数
	fmt.Printf("avg1=%v", avg1)
}
func getarrayavg(arr []float64, size int) float64 {
	var avg, sum float64
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	size1 := float64(size)
	avg = sum / size1
	return avg
}
func getintavg(arr []int, size int) int {
	var avg, sum int
	for i := 0; i < size; i++ {
		sum += arr[i]
	}
	avg = sum / size
	return avg
}
