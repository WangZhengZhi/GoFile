package main

import (
	"fmt"
	"os"
)

/*设备文件的使用*/
func main() {
	//os.Stdout.Close()    /*关闭后无法输出*/
	//fmt.Println("hello") /*往标准输出设备（屏幕）写内容*/
	/*标准设备文件，默认已经打开，用户可以直接使用，*/
	os.Stdout.WriteString("are u ok??\n")
	//os.Stdin.Close() /*关闭后无法输入*/
	var a int
	fmt.Println("请输入a:")
	fmt.Scan(&a)
	fmt.Printf("%d", a)

}
