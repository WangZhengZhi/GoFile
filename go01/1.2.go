package main

import (
	"fmt"
)

func main() {
	if true && true {
		fmt.Println("true")
	}
	if !false {
		fmt.Println("true")
	}
	for i := 0; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
J:
	for j := 0; j < 10; j++ {
		for k := 0; k < 10; k++ {
			if k > 5 {
				break J
			}
			println("此时的值为", k)
		}
	}
	for i := 0; i < 10; i++ {
		if i > 5 {
			continue
		}
		println(i)
	}
	list := []string{"a", "b", "c"}
	for k, v := range list {
		println(k, v)

	} //打印key value
	for pos, char := range "abc" {
		fmt.Printf("char=%c,pos=%d\n", char, pos)
	}
	var a int
	a = 100
	var b string
	b = "hello"
	fmt.Printf("a=%d,b=%s\n", a, b)
	switch a {
	case 0:
		fallthrough
	case 100:
		fmt.Printf("数值正确%d\n", 100)
	}

}

/*func Myfunc() {
	i := 0
Here:
	println(i)
	i++
	goto Here
}*/
//goto 语句
//Here  是一个标签
//goto Here即跳转
/*func Myfunc(){
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum)
} */
//for 循环语句
//for 循环有三种形式

//for init; condition;post{}类似C语言

//for condition{}类似while语句

//for{}死循环语句

//for循环也可以执行多个变量

//Reverse a
/*for i,j;len(a)-1;i<j;i,j=i+1,j-1{
	a[i],a[j]=a[j],a[i];//平行赋值
}*/
