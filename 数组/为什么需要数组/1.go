package main

import "fmt"

func main() {
	//数组，同一个类型的集合
	var id [50]int
	//操作数组，通过下标，从0开始到len（）-1
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d]=%d\n", i, id[i])
	}
}
