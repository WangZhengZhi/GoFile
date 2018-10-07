package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
切片做函数参数
 */
func initdata(slice []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(100)
	}

} /*初始化切片 给切片中的元素赋值为随机数字*/
func printSlice(slice []int) {
	for key, vaule := range slice {
		fmt.Printf("slice[%d]=%d\n", key, vaule)
	}
} /*打印slice*/
func sortSlice(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println("排序之后为：", slice)

} /*冒泡排序升序排列*/

func main() {
	n := 10
	slice := make([]int, n)
	initdata(slice)   //调用初始化函数
	printSlice(slice) //调用打印函数
	sortSlice(slice)  //调用冒泡排序升序排列函数

}
