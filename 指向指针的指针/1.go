package main

import "fmt"

//指向指针的指针
func main(){
var a int
var ptr *int
var pptr **int
a=30
ptr=&a
pptr=&ptr
fmt.Printf("a=%d\n",a)
fmt.Printf("*ptr=%d\n",*ptr)
fmt.Printf("**pptr=%d\n",**pptr)
}

