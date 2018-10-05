package main

/*
函数类型
*/
func main() {
	var result int
	result = Add(1, 2) //传统调用方法
	println(result)
	var Ftest FuncType //声明一个函数类型的变量，变量名字是Ftest
	Ftest = Add        //是变量就可以赋值
	result = Ftest(10, 20)
	println(result)

}
func Add(numa, numb int) int {
	return numb + numa

}
func Minus(numa, numb int) int {
	return numa - numb

}

//函数也是一种数据类型
//通过type给一个函数类型起名字
type FuncType func(int, int) int //没有函数名字。没有大括号
//FuncType是一个函数类型
