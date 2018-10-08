package main

import "fmt"

/*
方法值
*/
type Person struct {
	name string
	sex  byte
}

func (p Person) SetValueInfo() {
	fmt.Println("set value info")
}

func (p *Person) SetPointerInfo() {
	fmt.Println("set pointer info")
}

func main() {
	s := Person{"wang", 'M'}
	s.SetPointerInfo() /*传统的调用方式*/
	s.SetValueInfo()   /*传统的调用方式*/
	Vfunc := s.SetValueInfo
	Pfunc := s.SetPointerInfo
	Vfunc() /*方法值，调用函数的时候，无需再传递接收者，隐藏了接收者*/
	Pfunc() /*方法值，调用函数的时候，无需再传递接收者，隐藏了接收者*/
	/*等价与s.setvalueinfo s.setpointerinfo*/
	/*或者如下*/

}
