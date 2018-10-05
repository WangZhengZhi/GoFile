package main

import "fmt"

//匿名函数练习

func main() {
	/*返回一个*/
	func(numa, numb int) (result int) {
		if numa > numb {
			result = numa
		} else {
			result = numb
		}
		fmt.Println("max num is ", result)
		return result

	}(1, 2) //匿名函数的声明以及调用
	/*返回两个*/
	func(numa, numb int) (max, min int) {
		if numa > numb {
			max = numa
			min = numb
		} else {
			max = numb
			min = numa
		}
		fmt.Printf("max is :%d min is %d", max, min)
		return numa, numb
	}(10, 11)
}
