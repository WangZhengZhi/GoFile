package test

import "fmt"

type Student struct {
	Id   int
	Name string
	Sex  byte
}

func PrintStruct(s Student) {
	fmt.Println("s", s)

}

func myfunc() {
	fmt.Println("this is myfunc")
}
func Myfunc() {
	fmt.Println("this is Myfunc")
}
