package main

import "fmt"

/*
copy()函数的使用
 */
func main() {
	srcSlice := []int{1, 2, 3, 4}
	dstSlice := []int{5, 6, 7, 8, 9, 10}
	copy(dstSlice, srcSlice) //copy(目的切片，源切片)
	fmt.Println(dstSlice)    //输出为：1 2 3 4 9 10
	/*
	会根据目的切片做出裁剪
	把源切片的位置放到目的切片的位置
	 */

}
