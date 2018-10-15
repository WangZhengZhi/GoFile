package main

import (
	"fmt"
	"net"
)

/*tcp服务器问题*/
func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer listener.Close() //关闭

	//阻塞等待用户链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err:", err)
			continue //下次再来
		}
		//接收用户的请求
		buf := make([]byte, 1024) //1024大小缓冲区域
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("err1:", err)
			continue
		}
		fmt.Println("buf:", string(buf[:n])) //n是接受字节的长度
		defer conn.Close()                   //关闭当前用户链接
	}

}
