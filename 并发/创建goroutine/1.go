package main

import (
	"fmt"
	"time"
)

/*goroutine*/
func NewTask() {
	for {

		fmt.Println("this is new task")
		time.Sleep(time.Second * 1) /*延时1s*/

	}
}
func main() {
	for {

		fmt.Println("this is main goroutine")
		time.Sleep(time.Second * 1) /*延时1s*/
		//NewTask()                   /*无法调用*/
		go NewTask() /*新建一个协程,新建一个任务*/
	}
}
