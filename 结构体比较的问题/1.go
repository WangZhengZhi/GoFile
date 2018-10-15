package main

import "fmt"

func main() {
	s1 := struct {
		name string
		id   int
		//age int
	}{"hello", 18}
	s2 := struct {
		name string
		id   int
	}{"hello", 18}
	if s1 == s2 {
		fmt.Println("可以比较")
	}

}
