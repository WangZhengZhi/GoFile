package main

import (
	"calc"
	"fmt"
)
func init(){
	fmt.Println("这里是主函数的init函数")
}

func main() {
	a := calc.Add(10, 20)
	fmt.Println(a)
	b := calc.Minus(20, 10)
	fmt.Println(b)
}
