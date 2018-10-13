package main

func main() {
	//创建一个channel。双向的
	ch := make(chan int)
	//双向channel 可以隐式转化为单向channel
	var writech chan<- int = ch
	writech <- 666 //只能写的chan不能读
	//<-writech只能写不能读
	var readch <-chan int = ch
	<-readch //只能读不能写
	//readch <- 666只能读不能写

	//单项无法转化为双向
	//var ch2 chan int =readch
}
