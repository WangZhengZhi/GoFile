package main

import (
	"fmt"
)

func main() {
	var a, b, c chan int //只有channel  可以send /recived
	var i1, i2 int
	select {
	case i1 = <-c:
		fmt.Printf("%d recived from %d ", i1, c)
	case b <- i2:
		fmt.Printf("%d send to %d", i2, b)
	case i3, ok := (<-a):
		if ok {
			fmt.Printf("%d recived from a", i3)
		} else {
			fmt.Printf("%d is closed", a)
		}
	default:
		fmt.Println("NO communication")

	}
	/*
		每个case都必须是一个通信
		所有channel表达式都会被求值
		所有被发送的表达式都会被求值
		如果任意某个通信可以进行，它就执行；其他被忽略。
		如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
		否则：
		如果有default子句，则执行该语句。
		如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
	*/

}
