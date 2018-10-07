package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//新切片
	s1 := a[2:5] //从a[2]开始，取三个元素
	fmt.Println("s1", s1)
	s1[0] = 666
	fmt.Println("a:", a)
	s2 := s1[2:7]
	fmt.Println("s2", s2)
	s2[1] = 777
	fmt.Println("s2", s2)
	fmt.Println("a", a)

}
