package main

import (
	"fmt"
)

func main() {
	for _, value := range feiNum(5) {
		fmt.Printf("斐波那契数列为:%v\n", value)
	}
	arraynum := [...]int{1, 2, 3, 4, 5}
	sclicenum := arraynum[:]
	maxnum := max(sclicenum)
	fmt.Printf("maxnum=%v\n", maxnum)
	minnum := min(sclicenum)
	fmt.Printf("minnum=%v\n", minnum)
	/*str := "a"
	for i := 0; i < 4; i++ {
		fmt.Printf("%v\n", str)
		str = str + "a"
	}*/

}
func feiNum(num int) []int {
	array := make([]int, num)
	array[0] = 1
	array[1] = 1
	for n := 2; n < num; n++ {
		array[n] = array[n-1] + array[n-2]

	}
	return array
} //斐波那契函数
func max(slice []int) (maxNum int) {
	for _, value := range slice {
		if value > maxNum {
			maxNum = value

		}
	}
	return
} //筛选slice中最大值的函数

func min(slice []int) (minnum int) {
	for _, value := range slice {
		if value < minnum {
			minnum = value
		}
	}
	return

} //筛选slice中最小值的函数
