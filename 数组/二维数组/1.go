package main

import "fmt"

/*
二维数组
有多少个[]就有多少维
有多少[]就用多少循环
 */
func main() {
	k := 0
	var a [3][4]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])
		}
	}
	/*
	多维数组如何初始化
	 */
	b := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println("b:", b)
	/*
	部分初始化，没有初始化的值为0
	 */
	c := [3][4]int{
		{1, 2, 3, 4,},
	}
	fmt.Println("c:", c)
	/*
	部分初始化的情况二
	 */
	d := [3][4]int{
		{1, 2},
	}
	fmt.Println("d:", d)
	/*
	指定初始化
	 */
	e := [3][4]int{
	    	1:{1,2,3,4},
	}
	fmt.Println("e:",e)
}
