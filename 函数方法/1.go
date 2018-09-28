package main

import (
	"fmt"
)

//Go 语言中同时有函数和方法。
//一个方法就是一个包含了接受者的函数，
//接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
//所有给定类型的方法属于该类型的方法集
/*


func (variable_name variable_data_type) function_name() [return_type]{
   函数体
}
*/

func main() {
	var r1 Circle //定义一个Circle类型的数据
	r1.r = 10     //结构体中的数据定义值
	fmt.Println(r1.getArea())

}

type Circle struct {
	r float64 //圆的半径
} //数据结构

func (c Circle) getArea() float64 {
	return c.r * c.r * 3.14
}
