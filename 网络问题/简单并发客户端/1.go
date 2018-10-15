package main

import (
	"fmt"
	"net"
	"os"
)

/*简单的并发客户端*/
func main() {
	/*主动链接服务器*/
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	/*程序调用完毕关闭链接*/
	defer conn.Close()

	go func() {
		//从键盘输入内容给服务器发送内容
		str := make([]byte, 1024)
		for {

			n, err := os.Stdin.Read(str) //从键盘读取内容，放在str
			if err != nil {
				fmt.Println("os.Stdin serr:", err)
				return
			}
			//把输入的内容给服务器发送
			conn.Write(str[:n])
		}

	}()
	/*接收服务器回复的数据*/
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //接收服务器请求
		if err != nil {
			fmt.Println("conn.read err:", err)
			return
		}
		fmt.Println("buf:", string(buf[:n])) //  打印接收到的内容
	}
}
