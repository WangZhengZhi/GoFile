package main

import (
	"fmt"
)

/*
	slice appeend()
	向slice加入数据
*/
func main() {
	s0 := []int{0, 0}
	s1 := append(s0, 1)
	fmt.Println("s1=", s1)
	s2 := append(s1, 2)
	fmt.Println("s2=", s2)
	var a = []int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	n1 := copy(s, a[0:])
	fmt.Println(n1)
	fmt.Println(s)
	n2 := copy(s, s[2:])
	fmt.Println(n2)
	fmt.Println(s)
	/*
		此时的n 的值其实是你copy的个数
		而S定义时已经告诉赋值的位数
	*/
	/*
		map
		map定义的方式：
		map[<from type>]<to type>
		类似字典
	*/
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}
	fmt.Printf("%d\n", monthdays["Jan"])
	//判断一年有多少天
	year := 0
	for _, days := range monthdays {
		year += days
	} //不用数据月份只用天，所以用_省略
	fmt.Printf("%d\n", year)
	monthdays["unknowMonth"] = 40
	fmt.Printf("%d\n", monthdays["unknowMonth"]) //给MAp增加数据
	monthdays["Feb"] = 29                        //给MAP更改数据
	fmt.Printf("%d\n", monthdays["Feb"])
	key, ok := monthdays["Jan"]
	if ok {
		fmt.Println("Jan=", key)
	}
	//检查元素是否存在----简单写法
	//更加符合GO语言
	var value int
	var flag bool
	value, flag = monthdays["Nov"]
	if flag {
		fmt.Println("Nov=", value)
	}
	//检查元素是否存在 原始写法
	//更加符合逻辑
	delete(monthdays, "unknowMonth") //删除map中的数据
	fmt.Println(monthdays)
}
