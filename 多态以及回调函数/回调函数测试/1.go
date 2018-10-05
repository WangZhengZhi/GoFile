package main

//回调函数
type FuncType func(int, int) int

func Add(numa, numb int) int {
	return numb + numa

} //加法
func Miuns(numa, numb int) int {
	return numa - numb

} //减法
func Calc(numa, numb int, Ftest FuncType) (result int) {
	println("calc start!")
	result = Ftest(numa, numb)
	return result
}

func main() {
	a := Calc(1, 2, Add)
	println(a)
	b := Calc(1, 1, Miuns)
	println(b)
}
