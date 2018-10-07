package main

import "fmt"

/*
map[key_type]value_type
key_type是唯一的
map的基本使用
 */
func main() {
	//定义一个变量，类型为map[int]string
	var map1 map[int]string
	fmt.Println("map =", map1)
	//对于map只有len没有cap
	fmt.Println(len(map1))
	//可以通过make创建map
	map2 := make(map[int]string)
	fmt.Println("map2 =", map2)
	//可以通过make指定长度,只是指定了容量，里面是没有数据的
	map3 := make(map[int]string, 10) //make(map[int]string,len())
	fmt.Println("len(map3)=", len(map3))
	map3[1] = "make" //map元素的操作
	map3[2] = "go"
	map3[3] = "cpp"
	fmt.Println("map3:=", map3)
	fmt.Println("len(map3)	=", len(map3))
	/*
	key是唯一的
	每次增加元素，都是可以增加len的长度*/
	map4:=map[int]string{1:"make",2:"go",3:"cpp"}
	fmt.Println("map4:=",map4)
}
