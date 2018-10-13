package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

/*通过channel实现同步*/
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}

}
func Person1() {
	Printer("hello")
	ch <- 0
}
func Person2() {
	<-ch
	Printer("world")

}

func main() {
	go Person1()
	go Person2()
	for {
		fmt.Printf("")
	}
}
