package main

import "fmt"

/*
cap:容量
 */
func main() {
	a := []int{1, 2, 3, 4, 5}
	slice := a[0:3:5]
	fmt.Println("slice:", slice)          //1 2 3
	fmt.Println("len(slice)", len(slice)) //3-0
	fmt.Println("cap(slice)", cap(slice)) //5-0
	slice1 := a[1:4:5]
	fmt.Println("slice1:", slice1)          //2 3 4
	fmt.Println("len(slice1)", len(slice1)) //4-1
	fmt.Println("cap(slice1)", cap(slice1)) //5-1

}
