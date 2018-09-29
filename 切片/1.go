package main

import "fmt"

//go语言切片是对数组的抽象
//go数组的长度是不可以改变的
//所以引出切片的功能可以追加元素



/*切片的声明方式
1 ：slice:=make([]type,len)
切片初始化：
s:=[]int{1,2,3}初始化切片
直接初始化切片
slice:=arr[startIndex:endIndex]从下标startIndex ->enIndex-1的范围
*/
func main(){
var numbers=make([]int,3,5)//make([]type,len,cap)len为长度cap为测量切片的长度
printSlice(numbers)
var nilnum []int//尚未初始化的切片是nil
if nilnum==nil{fmt.Printf("切片是空的\n")}
printSlice(nilnum)
numbers1:=[]int{1,2,3,4,5}
printSlice(numbers1)
fmt.Printf("numbersa=%v\n",numbers1)
fmt.Printf("numbersa[1:4]=%v\n",numbers1[1:4])
fmt.Printf("numbersa[:]=%v\n",numbers1[:])
fmt.Printf("numbersa[1:]=%v\n",numbers1[1:])
fmt.Printf("numbersa[:4]=%v\n",numbers1[:4])
num2:=numbers1[:4]//切片中的cap（）问题，
// 假如：numx:=numy[:someNum] cap的值还是和numy一样的
//假如numx:=numy[someNum:]cap的值是:cap(numy)-someNum

printSlice(num2)


}
func printSlice(x[]int){
	fmt.Printf("len=%d,cap=%d,slice=%v\n",len(x),cap(x),x)
}


