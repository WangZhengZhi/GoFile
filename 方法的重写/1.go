package main

import "fmt"

/*
方法的继承
*/
type Person struct {
	name string
	sex  byte
	age  int
}

//person 类型实现了一个方法
func (tmp *Person) PrintInfo() {
	fmt.Printf("name=%s sex=%c age=%d", tmp.name, tmp.sex, tmp.age)
}

//有个学生继承了person 成员和方法都继承了
type Student struct {
	Person
	id   int
	addr string
}

//student也实现了一个方法，这个方法和person方法同名，这种方法叫重写
func (p *Student) PrintInfo() {
	fmt.Println("student:p", p)
}

func main() {
	s := Student{Person{"mike", 'M', 21}, 666, "BeiJing"}
	s.PrintInfo() //调用的是student中的方法
	/*也是就近原则，先找本作用域的方法，找不到，再用继承的方法*/
	/*不想调用student*/
	s.Person.PrintInfo() /*调用person中的方法*/

}
