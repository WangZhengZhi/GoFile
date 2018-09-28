package main

var a = 6
var b int

func main() {
	numa()
	numb()
	numa() //再次调用numa此时的a的值已经变为上一次函数调用时所赋值
	b = 1
	println(b)
	numf()
	for i := 0; i < 5; i++ {
		defer println(i) //defer函数延迟函数，按照后进先出的顺序执行
	}

}
func numa() {
	println(a)
}
func numb() {
	a = 5
	println(a)
}
func numf() {
	b = 0
	println(b)
	numg()

}
func numg() {
	println(b)
} //局部变量仅仅在执行定义她的函数时2有效
func f() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}
