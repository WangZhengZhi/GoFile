package main

import "fmt"

/*关闭channel*/
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //写入数据
		}
		close(ch)
	}()
	/*for {
		if num, ok := <-ch; ok == true { //读取数据
			fmt.Println("num:=", num)
		} else {
			break //chan关闭，跳出循环
		} //遍历内容
	}*/
	for num := range ch {
		fmt.Println("num=", num)
	} //迭代方式
}
