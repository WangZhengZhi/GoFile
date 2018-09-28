package main

import (
	"fmt"
)

func main() {
	addfunc := add(2, 4)
	fmt.Println(addfunc())
	fmt.Println(addfunc())
	fmt.Println(addfunc())
	fmt.Println(addfunc())
	fmt.Println(addfunc())
	fmt.Println(addfunc())
	somethingi := something(1, 2)
	fmt.Println(somethingi())
	fmt.Println(somethingi())
	fmt.Println(somethingi())
	fmt.Println(somethingi())
	fmt.Println(somethingi())
	fmt.Println(somethingi())
	fmt.Println(somethingi())
	fmt.Println(somethingi())

}
func add(x, y int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i += 1
		return i, x + y //此处决定匿名函数参数的返回类型和数量,同时也声明了函数参数的类型和数量
	}
} //带参数的闭包函数
func something(x, y int) func() int {
	var i int
	return func() int {
		i += 1
		return i
	}
}
