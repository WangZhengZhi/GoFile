package main

import "fmt"

/*接口的继承*/
type Humaner interface { /*子集*/
	SayHi()
}
type Personer interface { /*超集*/
	Humaner //匿名字段，继承了sayHI（）方法
	Sing(lrc string)
}
type Student struct {
	name string
	age  int
} /*Student实现了SayHi()*/

func (s *Student) SayHi() {
	fmt.Println("学生信息", s.name, s.age)
}
func (s *Student) Sing(lrc string) {
	fmt.Println("学生在唱:", lrc)
}

func main() {
	/*var i Personer*/
	/*s := &Student{"wang", 19}*/
	/*i = s*/
	/*i.SayHi() /*继承过来的方法*/
	/*i.Sing("yesterday once more")*/
	/*超集可以转化为子集*/
	var superi Personer /*超集*/
	superi = &Student{"wang", 19}
	var tinyi Humaner /*子集*/
	tinyi = superi    /*超集可以转换为子集,反过来不行*/
	tinyi.SayHi()

}
