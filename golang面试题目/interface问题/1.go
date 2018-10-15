package main

import "fmt"

/*interface()问题*/
func main() {
	var x *int = nil
	Foo(x)

}
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("空接口")
	} else {
		fmt.Println("非空接口")
	}

}
