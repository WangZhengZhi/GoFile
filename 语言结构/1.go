package main

import "fmt"

//go语言结构体
//type struct_vartiable_name struct{
// info 1
// info 2
// info 3}
type Books struct {
	title string
	author string
	subject string
	book_id int
}
func main(){
var book1 Books

book1.title="golang"
book1.author="google"
book1.subject="cs"
book1.book_id=1
fmt.Printf("title:%v\n",book1.title)
fmt.Printf("author:%v\n",book1.author)
fmt.Printf("subject:%v\n",book1.subject)
fmt.Printf("id:%v\n",book1.book_id)
var book2 Books
book2.title="go"
book2.author="tom"
book2.subject="CS"
book2.book_id=2
fmt.Printf("title:%v\n",book2.title)
fmt.Printf("author:%v\n",book2.author)
fmt.Printf("subject:%v\n",book2.subject)
fmt.Printf("id:%v\n",book2.book_id)
printbook(book1)
printbook1(&book2)//结构体指针函数调用


}
/*
结构体也可以作为函数的参数
 */
 func printbook(book Books){
	 fmt.Printf("title:%v\n",book.title)
	 fmt.Printf("author:%v\n",book.author)
	 fmt.Printf("subject:%v\n",book.subject)
	 fmt.Printf("id:%v\n",book.book_id)
 }
 func printbook1(book *Books){
	 fmt.Printf("title:%v\n",book.title)
	 fmt.Printf("author:%v\n",book.author)
	 fmt.Printf("subject:%v\n",book.subject)
	 fmt.Printf("id:%v\n",book.book_id)
 }// 结构体指针打印


