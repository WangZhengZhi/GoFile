package main

import "fmt"

/*
map 的遍历
 */
func main() {
	map1 := map[int]string{1: "go", 2: "cpp", 3: "java"}
	for key, value := range map1 {
		fmt.Printf("key=%d value=%s\n", key, value)
	} //第一个返回key第二个返回value,遍历的结果是无序列的
	/*
	如何判断一个key是否存在
	 */
	value, ok := map1[1]
	if ok == true {
		fmt.Println(value)
	} else {
		fmt.Println("不存在")
	}

}
