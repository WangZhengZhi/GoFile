package main

import "fmt"

/*
map做函数参数
引用传递
 */
func test(maps map[int]string) {
	delete(maps, 1)
} //删除某个key
func main() {
	map1 := map[int]string{1: "go", 2: "cpp", 3: "java"}
	test(map1)
	fmt.Println(map1)
}
