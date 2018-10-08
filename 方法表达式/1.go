package main

import "fmt"

/*
方法表达式
*/
type Person struct {
	name string
	age  int
}

func (p Person) SetValueInfo() {
	fmt.Println("SetValueInfo")
}
func (p *Person) SetPointerInfo() {
	fmt.Println("SetPointerInfo")
}

func main() {
	p := Person{"mike", 19}

	Pfunc := (*Person).SetPointerInfo
	Pfunc(&p) //等价于p.SetPointerInfo
	//显示的把接收者传递
	Vfunc := (Person).SetValueInfo
	Vfunc(p) //等价于p.SetValueInfo

}
