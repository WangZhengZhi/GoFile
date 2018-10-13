package main

/*字符串的操作*/
import (
	"fmt"
	"strings"
)

func main() {
	/*Trim*/
	str := "!!!hello!!!"
	afterstr := strings.Trim(str, "!")
	fmt.Println(afterstr) /*Trim的作用是去除头尾指定的字符*/

	/*Field*/
	str1 := "				are u ok			"
	afterstr1 := strings.Fields(str1)
	fmt.Println(afterstr1)      /*Field把string按照空格进行分隔,并且去掉多余的空格*/
	fmt.Println(len(afterstr1)) /*计算一句话有多少单词包含符号*/

	/*Contains*/
	str2 := "hello"
	fmt.Println(strings.Contains(str2, "he")) /*contains查看字符串是否包含某个字符串，包含返回true不包含返回false*/

	/*Joins*/
	slice := []string{"hello", "world"} /*切片才可以*/
	tmp := strings.Join(slice, "@")
	fmt.Println(tmp) /*指定切片中插入某个字符串*/

	/*Index*/
	fmt.Println(strings.Index("abcdef", "a")) /*存在则返回索引位置信息*/
	fmt.Println(strings.Index("abcdef", "b"))
	fmt.Println(strings.Index("abcdef", "go")) /*不存在则返回-1*/

	/*Repeat*/
	fmt.Println(strings.Repeat("go", 3)) /*会重复字符串，后一位是重复次数*/

	/*Split*/
	/*已指定的分隔符拆分*/
	fmt.Println(strings.Split("hello@world@!", "@")) /*已@拆分*/

}
