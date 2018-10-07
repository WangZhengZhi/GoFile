package main

import "fmt"

/*
值语义和引用语义的不同
*/
type Person struct {
	name string
	age  int
}

func (value Person) ValueInfo(name string, age int) {
	value.name = name
	value.age = age
	fmt.Printf("%p\n", &value) //打印指针地址
	fmt.Println("ValueInfo:", value)

} //值传递
func (pointer *Person) PointerInfo(name string, age int) {
	pointer.name = name
	pointer.age = age
	fmt.Printf("%p\n", pointer) //打印指针地址
	fmt.Println("ValueInfo:", pointer)

} //引用传递（指针）

func main() {
	/*
		值传递测试
	*/
	s := Person{"wang", 18}
	s.ValueInfo("li", 19)
	fmt.Println("main:", s)
	fmt.Printf("%p\n", &s) //打印指针地址
	/*
		引用传递测试

	*/
	s1 := Person{"wang", 19}
	(&s1).PointerInfo("zhao", 22)
	fmt.Println("main:", s1)
	fmt.Printf("%p\n", &s1) //打印指针地址

}
