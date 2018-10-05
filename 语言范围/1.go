package main

import "fmt"

//go语言范围

func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	} //使用range 求切片的和
	fmt.Printf("sum=%d\n", sum)
	//有的情况下需要知道索引值
	for key, num := range nums {
		if num == 3 {
			fmt.Printf("%d的索引是%d\n", num, key)
		}
	}
	//range 有时候也可以用在map上
	maps := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range maps {
		fmt.Printf("%v:%v\n", key, value)
	}
	//range也可以枚举
	for i, c := range "hello" {
		fmt.Printf("%v,%c\n", i, c) //输出unicode编码值
		//%c 输出字符，%d或者%v是输出编码值
		//第一个值是索引，第二个是字符（unicode编码值）

	}

}
