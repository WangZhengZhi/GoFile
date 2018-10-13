package main

import (
	"fmt"
	"strconv"
)

/*Append Format*/
/*Append系列函数，将其他类型数据转化为字符串，添加到现有的字节数组之中*/
func main() {
	/*转化为字符串后追加到字节数组*/
	slice := make([]byte, 0, 1024)
	/*给bool类型追加数据*/
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 1234, 10) /*追加1234，并且是已10进制方式追加*/
	slice = strconv.AppendQuote(slice, "abcdef")
	fmt.Println(string(slice)) /*转化为字符串再打印*/
	/*其他类型转化为字符串*/
	var str string
	str = strconv.FormatBool(false)
	str = strconv.FormatFloat(3.14, 'f', -1, 64)

	/*f代表小数点这仲格式，b 代表指数二进制方式，e/E是十进制指数
	prec=-1,尽量用小数点最少数量的，但又是必须的数字表示f
	f是float
	64是float64*/
	/*整型----->字符串*/
	str = strconv.Itoa(6666)

	str = strconv.FormatInt(98, 10) /*10进制去转化*/
	str = strconv.FormatInt(16, 2)  /*2进制*/
	fmt.Println(str)

	/*字符串转化为其他类型*/
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag:", flag)
	} else {
		fmt.Println(err) /*如果错误则会引发erro*/
	}
	/*把字符串------->整型*/
	num, err := strconv.Atoi("4556777")
	if err == nil {
		fmt.Println("num:", num)
	} else {
		fmt.Println("err:", err)
	}
}
