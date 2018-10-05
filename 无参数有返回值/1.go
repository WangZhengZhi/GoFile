package main

import "fmt"

func main() {
	num := myfunc()
	fmt.Printf("num=%d\n", num)
	num01 := myfunc01()
	fmt.Println("num01=", num01)
	numa, numb := myfunc02()
	fmt.Printf("numa=%d,numb=%d\n", numa, numb)
	max := funcmax(1, 2)
	fmt.Printf("max=%d\n", max)
	min := funcmin(1, 2)
	fmt.Printf("min=%d\n", min)
	maxnum, minnum := funcMAXMIN(1, 2)
	fmt.Printf("maxnum=%d minnum=%d\n", maxnum, minnum)
	//也可以通过匿名变量丢弃返回值
	/*
		maxnum,_:=funcMAXMIN(1,2)
		fmt.Printf("maxnum=%d \n", maxnum)
	*/

}
func myfunc() int {
	return 666
}

//无参数有返回值的函数
//有返回值的函数需要通过return返回
func myfunc01() (result int) {
	result = 888
	return result
} //最好给返回值一个类型
/*
返回多个值的函数
*/
func myfunc02() (numa, numb int) {
	numa, numb = 1, 2
	return numa, numb
}
func funcmax(numa, numb int) (max int) {
	if numa > numb {
		max = numa
	} else {
		max = numb
	}
	return max
}
func funcmin(numa, numb int) (min int) {
	if numa > numb {
		min = numb
	} else {
		min = numa
	}
	return min

}
func funcMAXMIN(numa, numb int) (max, min int) {
	if numa > numb {
		max = numa
		min = numb

	}
	if numa < numb {
		min = numa
		max = numb
	}
	return max, min

}
