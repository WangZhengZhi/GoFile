package main

import "fmt"

/*
结构体类型添加方法
*/
type Person struct {
	name string
	sex  byte
	age  int
}

func (temp Person) PrintInfo() {
	fmt.Println("temp=", temp)
} //方法，打印结构体数据
//通过一个函数给成员赋值
func (temp *Person) SetInfo(name string, sex byte, age int) {
	temp.name = name
	temp.sex = sex
	temp.age = age
}
func main() {
	p := Person{"age", 'M', 19}
	p.PrintInfo()
	var p2 Person
	(&p2).SetInfo("yoyo", 'G', 18)
	//(&p2).PrintInfo()
	//也可以使用这个'
	p2.PrintInfo()
}
