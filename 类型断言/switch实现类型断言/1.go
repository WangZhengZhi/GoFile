package main

import (
	"fmt"
)

/*switch实现类型断言*/
type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1                    //int
	i[1] = "hello"              //string
	i[2] = Student{"hello", 88} //student
	for index, data := range i {
		switch data.(type) {
		case int:
			fmt.Println("int:", index, data)
		case string:
			fmt.Println("string:", index, data)
		case Student:
			fmt.Println("Student:", index, data)
		}

	}
	for index, data := range i {
		switch value := data.(type) {
		case int:
			fmt.Println("int", index, value)
		case string:
			fmt.Println("string", index, value)
		case Student:
			fmt.Println("student:", index, value)
		}
	}
}
