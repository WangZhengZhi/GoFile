package main

import "fmt"

type Humaner interface {
	SayHi()
}
type Personer interface {
	Humaner /*继承sayHi()*/
	Sing(lrc string)
}
type Student struct {
	name string
	age  int
}

func (s *Student) SayHi() {
	fmt.Println("学生信息：", s.name, s.age)
}
func (s *Student) Sing(lrc string) {
	fmt.Println("学生在唱：", lrc)
}

func main() {
	var i Personer
	s := &Student{"wang", 21}
	i = s
	i.SayHi()
	i.Sing("yesterday once more")
}
