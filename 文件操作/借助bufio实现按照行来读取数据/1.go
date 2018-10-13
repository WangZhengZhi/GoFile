package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*借助bufio按照行来读取文件*/
func ReadFileLine(path string) {
	//打开文件
	data, err := os.Open(path)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//关闭文件
	defer data.Close()
	//建立缓冲区域，把内容先放在缓冲区域
	r := bufio.NewReader(data)
	//按照行来读取
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break //文件结束
			}
			fmt.Println("err:", err)
		}
		fmt.Println("buf", string(buf))

	}

}
func main() {
	var path string
	path = `C:\Users\apple\go\src\文件操作\文件的读写\demo.txt`
	ReadFileLine(path)
}
