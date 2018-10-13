package main

import (
	"fmt"
	"runtime"
)

/*gosched的使用*/
func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子协程")
		}
	}()
	for i := 0; i < 2; i++ {
		/*需要让出一个时间片，来执行子协程*/
		runtime.Gosched() //让子协程先执行
		fmt.Println("主协程")
	}
}
