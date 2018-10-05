package main

/*
通过递归实现累加
1+2+3+........+100
*/
func main() {

	var sum int
	sum = Add(100)
	println(sum)
	println(Addreverse(1)) //也是递归，反过来递归
}
func Add(num int) (sum int) {
	if num == 1 {
		sum = 0
		return 1
	}

	return num + Add(num-1)

}
func Addreverse(num int) (sum int) {
	if num == 100 {
		return 100
	}
	return num + Addreverse(num+1)
}
