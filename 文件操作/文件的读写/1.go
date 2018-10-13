package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	//打开文件，新建文件
	data, err := os.Create(path)
	if err != nil {
		fmt.Println("err:=", err)
		return
	}
	//使用完毕，关闭文件夹：
	defer data.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("%d\n", i) //i 的数据存储在buf中
		_, err1 := data.WriteString(buf)
		if err1 != nil {
			fmt.Println("err:", err1)
			return
		}

	}
}
func ReadFile(path string) {
	//打开文件
	data, err := os.Open(path)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//关闭文件
	defer data.Close()
	buf := make([]byte, 1024*2)
	//n从文件读取内容的长度
	n, err1 := data.Read(buf)
	//io.EOF文件的结尾
	if err1 != nil && err1 != io.EOF {
		fmt.Println("err1:=", err1)
		return
	}
	fmt.Println("buf:=", string(buf[:n]))

}

//每次读取一行

func main() {
	path := `C:\Users\apple\go\src\文件操作\文件的读写\demo.txt`
	//WriteFile(path)
	ReadFile(path)
}
