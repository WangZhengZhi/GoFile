package main

import (
	"fmt"
)

type stack struct {
	i    int
	data [10]int
}

//栈的定义
func main() {
	var s stack
	s.push(25)
	s.push(14)
	fmt.Printf("stack %v\n", s)
}
func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++

}
func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

//push pop出栈入栈
