package main

import "fmt"

//int/float64直接除会出现错误
//float64(int)/float64
func main() {
	var i int
	i = 10
	var j float64
	j = 100
	fmt.Println(j / float64(i))
}
