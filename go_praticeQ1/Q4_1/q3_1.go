package main

import (
	"fmt"
)

func main() {
	str := "asSASA ddd dsjkdsjs dk"
	s := []rune(str) //将S化为一个数字
	copy(s[4:], []rune("abc"))
	fmt.Println("before", str)
	fmt.Println("after", string(s)) //将S 转化为 string类型

}
