package main

import (
	"fmt"
	"time"
)

/*多任务资源竞争问题 */
func Printer(str string) {
	for _, value := range str {
		fmt.Printf("%c", value)
		time.Sleep(time.Second * 1)
	}
	fmt.Printf("\n ")
} //打印机
func Person1() {
	Printer("hello")

} //人物1
func Person2() {
	Printer("world")

} //人物2
func main() {
	/*建立2个协程*/
	go Person1()
	go Person2()
	for {
		fmt.Printf("")
	}
}
