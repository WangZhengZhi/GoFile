package main

import "fmt"

/*类型断言，if实现*/
type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1                  //int
	i[1] = "hello"            //string
	i[2] = Student{"mike", 8} //student
	/*第一个返回下标，第二个返回下标对应的值*/
	for index, data := range i {
		/*第一个是返回值，第二个是判断数据的真假*/
		if value, ok := data.(int); ok == true {
			fmt.Println("int", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Println("string", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Println("student", index, value)
		}
	}
}
