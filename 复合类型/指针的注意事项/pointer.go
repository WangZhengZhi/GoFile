package main

//不要操作没有合法指向的内存
import "fmt"

func main() {
	var p *int
	fmt.Println("p=", p)
	//*p=666
	//指针没有合法指项
	var a int
	p = &a
	*p = 666
	fmt.Println("*p=", *p)
}
