package main

import "fmt"

/*求未知&& || 的优先级别*/
func main() {
	var a bool = false
	var b bool = false
	var c bool = true
	if a && b || c {
		fmt.Println("&&优先")
	} else {
		fmt.Println("||优先")
	}
}
