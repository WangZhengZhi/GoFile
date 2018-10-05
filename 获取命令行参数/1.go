package main

import (
	"fmt"
	"os"
)

//获取命令行参数
func main() {
	list := os.Args //[]string类型/字符串
	n := len(list)
	fmt.Println("n=", n)

	for key, value := range list {
		fmt.Printf("key=%d value=%s\n", key, value)
	}
}
