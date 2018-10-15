package main

import (
	"fmt"
	"net"
	"strings"
)

/*简单的并发服务器*/
func HandleConn(conn net.Conn) {

	addr := conn.RemoteAddr().String() //获取客户端的网络地址信息
	fmt.Println(addr, "	addr connect successful")
	//读取用户数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	if "exit" == string(buf[:n-1]) {
		fmt.Println(addr, "程序退出")
		return
	} //客户端输入"exit"则退出
	fmt.Println("len(string(buf[:n]))", len(string(buf[:n])))
	//接收数据如果多一个所以减去一个：n-1
	fmt.Printf("%s	发送的数据给 	%s\n", addr, string(buf[:n]))

	conn.Write([]byte(strings.ToUpper(string(buf[:n])))) //数据转为大写再给用户发送

	defer conn.Close() //函数调用完毕关闭链接

} //处理用户请求
func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err:", nil)
		return
	}
	defer listener.Close() //关闭监听
	//接收多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		//处理用户发送的数据
		go HandleConn(conn) //新建一个协程

	}
}
