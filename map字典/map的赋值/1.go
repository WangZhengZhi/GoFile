package main

import "fmt"

/*
map的赋值
 */
func main() {
	map1 := map[int]string{1: "mike", 2: "yoyo"}
	//赋值，如果已经存在的key值，修改内容
	map1[1] = "cpp"
	fmt.Println("map1", map1)
	map1[3] = "go"
	fmt.Println("map1:", map1)
	/*
	追加，map底层自动扩容。和append类似
	 */

}
