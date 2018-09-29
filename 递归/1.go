package main

import "fmt"

//递归就是自己调用自己
/*
func recursion(){
recursion()}
func main(){
recursion}
golang支持递归，但是要设置退出条件，否则就一直执行
会引起错误

 */
func main() {
/*
阶乘
 */
var i int=15
fmt.Printf("%d的阶乘是%d\n",i,factorial(i))
for n:=0;n<10;n++{
	fmt.Printf("%d\t",feibo(n))
}
}
func factorial(n int)(result int){
	if(n>0){
		result=n*factorial(n-1)
		return result
	}
	return 1
}
func feibo(n int)int{
	if(n<2){
		return n
	}
	return feibo(n-1)+feibo(n-2)
}
