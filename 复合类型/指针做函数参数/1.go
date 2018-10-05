package main

import "fmt"

//地址传递
func Swap(a, b *int) {
	*a, *b = *b, *a //交换指针
}
func main() {
	a, b := 10, 20
	Swap(&a, &b)
	fmt.Println(a, b)
}
