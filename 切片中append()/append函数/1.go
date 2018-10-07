package main

import "fmt"

/*
append()函数
append（）扩容的特点
 */

func main() {
	s := []int{}
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 4) //向末尾增加元素
	fmt.Printf("s=%v len(s)=%d cap(s)=%d\n", s, len(s), cap(s))
	/*
	append()函数会智能的去控制底层数组的增长，一旦超出底层数组容量，通常以2倍容量重新分配底层数组，并复制原来的数据
	 */
	slice := make([]int, 0, 1) //len(slice)=0 cap(slice)=1
	oldCap := cap(slice)
	for i := 0; i < 8; i++ {
		slice = append(slice, i)
		if newCap := cap(slice); oldCap < newCap {
			fmt.Printf("oldcap=%d------> newcap=%d\n", oldCap, newCap)
			oldCap = newCap
		}
	}

}
