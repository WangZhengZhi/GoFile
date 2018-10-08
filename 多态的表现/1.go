package main

import "fmt"

/*多态的表现*/
/*定义一个普通函数，函数的参数为接口类型
只有一个函数却可以有不同的表现，这就是多态*/
type Humaner interface {
	SayHi()
}
type Person struct {
	name string
	age  int
}

func (P *Person) SayHi() {
	fmt.Println("say hi", P.name, P.age)
}

type Student struct {
	Person
	addr string
}

func (S *Student) SayHi() {
	fmt.Println("say hi", S.name, S.age, S.addr)
}

type Mystr string

func (mystr Mystr) SayHi() {
	fmt.Println("say hi", mystr)
}

func WhoIsSayHi(i Humaner) {
	i.SayHi()
}

func main() {
	P := &Person{"wang", 19}
	S := &Student{Person{"li", 20}, "bj"}
	var mystr Mystr = "zhao"
	/*多态的调用，一个函数有不同的表现*/
	WhoIsSayHi(P)
	WhoIsSayHi(S)
	WhoIsSayHi(&mystr)
	/*也可以通过slice实现，更加直观一点*/
	slice := make([]Humaner, 3) //类型为Humer CAP =3
	slice[0] = P
	slice[1] = S
	slice[2] = &mystr
	for _, value := range slice {
		WhoIsSayHi(value)
	} /*迭代的方式去实现输出更加方便*/

}
