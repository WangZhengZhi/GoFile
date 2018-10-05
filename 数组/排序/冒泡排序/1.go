package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a [5]int
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
	//fmt.Println(a)
	/*
	冒泡排序
	升序排列
	 */
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("升序排列：a[%d]=%d\n", i, a[i])
	}
	/*
	冒泡排序
	降序排列
	 */
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] < a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("降序排列：a[%d]=%d\n", i, a[i])
	}
}
