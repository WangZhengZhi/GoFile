package main

import (
	"fmt"
)

//编写一个函数，返回两个参数中正确的数字顺序
//f(2,7)->2,7
//f(7,2)->2,7
func main() {

	compare(2, 7)
	compare(7, 2)
	manyNumbers(1, 3, 4, 5, 6, 7, 8, 8)
	for _, value := range feiNum(3) {
		fmt.Printf("斐波那契数列为：%v\n", value)
	}

}
func compare(numa int, numb int) (int, int) {
	if numa > numb {
		fmt.Printf("%d,%d\n", numb, numa)
		return numb, numa
	} else {
		fmt.Printf("%d,%d\n", numa, numb)
		return numa, numb
	}
}

//变参数问题
//可以向函数注入许多类型一样的数据
func manyNumbers(nums ...int) {
	for _, v := range nums {
		fmt.Printf("%d\n", v)
	}
}

//斐波那契数列
//Xn=X(n-1)+X(n-2)
func feiNum(num int) []int {
	x := make([]int, num) //array的创建
	x[0], x[1] = 1, 1
	for n := 2; n < num; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x
}
