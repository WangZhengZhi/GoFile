package main

import (
	"fmt"
	"io"
	"os"
)

/*拷贝文件*/
func main() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("usage:xxx srcFile dstFile")
		return
	}
	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件不能和目的文件相同")
		return
	}
	//只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err1:", err1)
		return
	}
	//新建目的文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2:", err2)
		return
	}
	//操作完毕，关闭文件夹
	defer sF.Close()
	defer dF.Close()
	//核心处理，从源文件读取内容，写向目的文件，多少写多少
	buf := make([]byte, 4*1024)
	for {
		n, err := sF.Read(buf)
		if err != nil {
			fmt.Println("err:", err)
			if err == io.EOF {
				break //文件读取完毕
			}
			fmt.Println("err:", err)

		}
		//写入目的文件，多少写多少
		dF.Write(buf[:n])
	}
}
