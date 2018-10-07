package main

import "fmt"

/*
map的删除
 */
func main() {
	map3 := map[int]string{1: "mike", 2: "go", 3: "cpp"}
	fmt.Println("before delete",map3)
	delete(map3,1)
	fmt.Println("after delete",map3)
}
