package main

import "fmt"

/*
切片的创建方式
 */
func main() {
	/*
	自动推导类型
	 */
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("s:", s)
	/*
	make函数 make(slice_type,len,cap)
	 */
	s1 := make([]int, 5, 10)
	fmt.Printf("s1:%d len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	/*
	make(slice_type,len)不写cap。默认cap=len
	 */
	s3 := make([]int, 5)
	fmt.Printf("s3:%d len(s3):%d cap(s3):%d\n", s3, len(s3), cap(s3))


}
