package main

import "fmt"

/*接口的定义和实现
接口只关心行为*/
/*定义接口类型*/
type Humaner interface {
	//方法只有声明没有实现,有自定义类型（别的类型）实现
	SayHi()
}
type Student struct {
	name string
	id   int
}

/*student实现了方法*/
func (tmp *Student) SayHi() {
	fmt.Printf("%s %d ,say hi\n", tmp.name, tmp.id)

}

type Teacher struct {
	addr  string
	group string
}

/*teacher实现了方法*/
func (tmp *Teacher) SayHi() {
	fmt.Printf("%s %s ,say hi\n", tmp.addr, tmp.group)
}

/**/
type Mystr string

func (tmp *Mystr) SayHi() {
	fmt.Printf("Mystr[%s] say hi", *tmp)
}

func main() {
	/*定义接口类型的变量*/
	var i Humaner
	/*只是是实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值*/
	s := &Student{"mike", 333}
	i = s
	i.SayHi()
	t := &Teacher{"BJ", "go"}
	i = t
	i.SayHi()
	var str Mystr = "mike"
	i = &str
	i.SayHi()

}
