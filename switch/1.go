package main

import (
	"fmt"
)

//switch的用法
//switch后面的值可以是同类型的任意值
func main() {
	var rank string
	var flag int = 90
	switch flag {
	case 90, 100:
		rank = "A"
	case 80:
		rank = "B"
	case 70, 60:
		rank = "C"

	default:
		rank = "F"
	}
	switch {
	case rank == "A":
		fmt.Println("优秀")
	case rank == "B":
		fmt.Println("良好")
	case rank == "C":
		fmt.Println("一般")
	default:
		fmt.Println("一般")
	}
	fmt.Printf("你的等级是%d\n", flag)
	var score float32 = 78.0
	switch score {
	case 10.0:
		fmt.Println("成绩差")
	case 78.0:
		fmt.Println("及格")
	default:
		fmt.Println("不符合")
	} //switch 后面的选择值可以是任意相同的类型
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println("nil")
		fmt.Printf("%T", i)
	case int:
		fmt.Println("int")
		fmt.Printf("%T", i)
	case float32:
		fmt.Println("float32")
		fmt.Printf("%T", i)
	default:
		fmt.Println("unknow type!")

	} //switch 判断interface中变量实际存储的类型

}
