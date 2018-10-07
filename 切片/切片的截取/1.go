package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(cap(array))
	s1 := array[:6]
	/*
	len(s1)=6-0
	cap(s1)=cap(array)-0
	 */
	fmt.Println(s1, len(s1), cap(s1))
	s2 := array[1:3]
	/*
	len(s2)=3-1
	cap(s2)=cap(array)-1
	 */
	fmt.Println(s2, len(s2), cap(s2))
	s3 := array[1:3:4]
	fmt.Println(cap(s3), len(s3))
	/*
	cap(s3)=4-1
	len(s3)=3-1
	 */
	s4 := array[1:]
	fmt.Println(len(s4), cap(s4))
	/*
	len(s4)=len(array)-1=9-1
	cap(s4)=cap(s4)-1=9-1
	 */
}
