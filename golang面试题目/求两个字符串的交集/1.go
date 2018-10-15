package main

import "fmt"

func main() {
	str1 := "abcdefghj"
	str2 := "hghcdeshgv"
	for _, value := range str1 {
		for _, value1 := range str2 {
			if value == value1 {
				fmt.Printf("相同的交集为%c：", value)
			}
		}
	}
}
