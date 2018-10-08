package main

import "fmt"

/*普通变量的方法集*/
type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetValueInfo() {
	fmt.Println("set value info ")
}
func (p *Person) SetPointerInfo() {
	fmt.Println("set pointer info")
}

func main() {
	p := Person{"wang", 'M', 19}
	p.SetValueInfo()
	p.SetPointerInfo()
	/*指针这样的形式不用手动转化
	内部使用（&p）.setPointerInfo*/

}
