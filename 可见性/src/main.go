package main

import "可见性/src/test"

func main() {
	test.Myfunc() //package.func_name
	s := test.Student{1, "wang", 'M'}
	test.PrintStruct(s)
}
