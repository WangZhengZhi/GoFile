package main

import "fmt"

type User struct {
}
type Myuser1 User
type Myuser2 = User

func (i Myuser1) m1() {
	fmt.Println("myuser1.m1")
}
func (i User) m2() {
	fmt.Println("user.m2")
}
func main() {
	var i1 Myuser1
	var i2 Myuser2
	i1.m1()
	i2.m2()

}
