package main

import (
	"fmt"
)

//逆转字符串
func main() {
	/*s := "abcd"
	a := []rune(s)
	for i,j:=0,len(a)-1;i<j;i,j=i+1,j-1{
		a[i],a[j]=a[j],a[i]
	}
	fmt.Printf("%s\n",string(a))*/
	s := "abcd"
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i] //每一行使用多个变量只能使用平行赋值
	}
	fmt.Printf("%s\n", string(a))
	PrintSome(5)
	rec(5)

}
func PrintSome(in int) int {

	fmt.Printf("%d\n", in)
	return in

}
func rec(i int) {
	if i == 10 {
		return
	}
	rec(i + 1)
	fmt.Printf("%d\n", i)
} //递归函数
