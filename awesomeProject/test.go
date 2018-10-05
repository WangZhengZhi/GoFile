package main

import "fmt"

func main() {
	var pwa int
	//var pwb int
	//fmt.Println("输入密码")
	fmt.Scan(&pwa)
	if pwa == 5211314 {
		fmt.Println("密码正确")
	} else {
		fmt.Println("密码错误")
	}

}
