package main

import (
	"fmt"
	"unicode/utf8"
)

/*
01建立一个程序打印
A
AA
AAA
AAAA
AAAAA
AAAAAA
.............
*/
func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Println("")
	}
	// AAAAAAAAAA
	// AAAAAAAAA
	// AAAAAAAA
	// AAAAAAA
	// AAAAAA
	// AAAAA
	// AAAA
	// AAA
	// AA
	// A
	//这样的形式 需要将 j<i改为 j<10-i
	/*
		02建立一个程序统计字符串里的字符数量
		并输出字节数

	*/
	str := "ass SASA dddd sjkdsjs dk"
	fmt.Printf("String %s\nLength:%d,Runes:%d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
	//len([]byte(str))可以计算出字符数量
	//utf8.Runcount([]byte(str))可以计算出字节数

	/*
		03替换位置4开始的三个字符为abc

	*/
	c := []rune(str)
	copy(c[4:7], []rune("abc"))
	fmt.Printf("before:%s\n", str)
	fmt.Printf("after:%s\n", string(c))

}
