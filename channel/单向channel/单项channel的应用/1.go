package main

import "fmt"

/*单向channel 的应用*/
func productor(out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i
	}
	close(out) //写入完毕，要关闭channel

} //写入数据,不能读取
/*

 */
func costumer(in <-chan int) {
	for num := range in {
		fmt.Println("num=", num)
	}
} //读取数据，不能写
func main() {
	//创建一个双向通道
	ch := make(chan int)
	go productor(ch) //开一个子协程
	costumer(ch)

}
