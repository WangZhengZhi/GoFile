package main

import (
	"fmt"
)

func main() {
	var array [10]int
	array[0] = 42
	array[1] = 13
	fmt.Printf("array[0]=%d,array[1]=%d\n", array[0], array[1])
	var numA [3]int //声明未初始化的数组
	for i := 0; i < 3; i++ {
		fmt.Printf("num[%d]=%d\n", i, numA[i])
	}
	//快速声明 赋值的数组
	numB := [...]int{1, 2, 3} //也可以写作 numB:=[3]int{1,2,3}
	//numB := [3]int{1, 2, 3}
	for i := 0; i < 3; i++ {
		fmt.Printf("numB[%d]=%d\n", i, numB[i])
	}
	//多维数组声明
	//方法一
	a := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d]=%d\n", i, j, a[i][j])
		}
	}
	//很久之前的版本
	//方法二
	b := [3][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("b[%d][%d]=%d\n", i, j, b[i][j])
		}
	}
	//方法三
	c := [3][2]int{[...]int{1, 2}, [...]int{3, 4}, [...]int{5, 6}}
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("c[%d][%d]=%d\n", i, j, c[i][j])
		}
	}
	arrayList := [...]int{1, 2, 3, 4, 5}
	s1 := arrayList[2:4] //序列号 2到3
	s2 := arrayList[1:5] //1-4
	s3 := arrayList[:]   //全部
	s4 := arrayList[:4]  //0-4，可以简写
	s5 := s2[:]

	fmt.Println("s1=!!!!!!", s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	//slice超出 array索引的情况
	var arrayNum1 [100]int
	slice := arrayNum1[:99]
	slice[98] = 10
	fmt.Println(slice[98])
	/*
		slice[99]=100
		fmt.Println(slice[99])
		会出现报错
		超出索引范围
	*/
	/*
		slice append()
		slice copy()
	*/
	

}
