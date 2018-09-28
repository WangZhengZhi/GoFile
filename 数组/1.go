package main

import (
	"fmt"
)

func main() {
	var array [10]int
	for i := 0; i < 10; i++ {
		array[i] = i
	}
	for j := 0; j < 10; j++ {
		fmt.Printf("%d,%d\n", j, array[j])
	}
	/*二维数组*/
	a := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	} //定义二维数组
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%d\n", a[i][j])
		}
	} //输出二维数组

}

//数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，
//这种类型可以是任意的原始类型例如整形、字符串或者自定义类型。

//相对于去声明number0, number1, ..., and number99的变量，
//使用数组形式numbers[0], numbers[1] ..., numbers[99]更加方便且易于扩展。

//数组元素可以通过索引（位置）来读取（或者修改），
//索引从0开始，第一个元素索引为 0，第二个索引为 1，以此类推。
//多维数组
