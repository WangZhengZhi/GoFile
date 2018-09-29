package main

import "fmt"

//append() copy()函数
func main(){
var numbers []int
printSlice(numbers)
/*向切片增加一个元素*/
numbers=append(numbers,1)
printSlice(numbers)//
/*向切片增加多个元素*/
numbers=append(numbers,2,3)
printSlice(numbers)
/*创建切片增加切片容量*/
numbers1:=make([]int,len(numbers),cap(numbers)*2)//增加切片容量是numbers的2倍
/*copy numbers的内容到 numbers1*/
copy(numbers1,numbers)//copy(目标切片，被copy的切片)
printSlice(numbers1)

}
func printSlice(x[]int){
	fmt.Printf("len=%d,cap=%d,x=%v\n",len(x),cap(x),x)
}


