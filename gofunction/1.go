package main

import (
	"fmt"
)

//golang function
func main() {
	var i int
	var j int
	i = 10
	j = 19
	fmt.Printf("max=%d\n", max(i, j))
	fmt.Printf("min=%d\n", min(i, j))
	i1, j1 := swapnum(i, j)       //10 ,19
	fmt.Printf("%d,%d\n", i1, j1) //19,10
	str1 := "hello"
	str2 := "world"
	str11, str22 := swapstring(str1, str2)
	fmt.Println(str11, str22)

}
func max(numa int, numb int) int {
	if numa > numb {
		return numa
	} else {
		return numb
	}
} //返回最大值函数
func min(numa int, numb int) int {
	if numa > numb {
		return numb
	} else {
		return numa
	}
} //返回最小值函数
func swapnum(numa int, numb int) (int, int) {
	return numb, numa
} //交换int类型数据
func swapstring(str1, str2 string) (string, string) {
	return str2, str1
} //交换string类型数据
