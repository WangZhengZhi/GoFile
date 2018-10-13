package main

import (
	"fmt"
	"regexp" /*正则表达式的包*/
)

/*正则表达式强大但是效率低下*/
/*正则表达式*/
func main() {
	//1>解释规则
	array := "abc abd ab1 ab2 ab3 abz tab"
	reg1 := regexp.MustCompile(`ab.`)
	if reg1 == nil {
		fmt.Println("wrong!")
		return
	}

	//2>根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(array, -1)
	fmt.Println(result1)

	str := "Hello 你好 Golang"
	reg2 := regexp.MustCompile(`[a-z]+`) /*匹配连续的小写字母*/
	fmt.Println(reg2.FindAllStringSubmatch(str, -1))

	reg3 := regexp.MustCompile(`[\p{Han}]+`) /*查找连续的汉字 */
	fmt.Println(reg3.FindAllStringSubmatch(str, -1))

	str1 := "abc a7c a9c a9b"
	reg4 := regexp.MustCompile(`a\dc`)
	fmt.Println(reg4.FindAllStringSubmatch(str1, -1)) /*查找a[0-9]c*/

	/*查找一个array中的小数*/
	/*`+`匹配前一个字符一次或者多次*/
	str2 := "1.34 567 hihihi 7.18 8.19 9.99"
	reg5 := regexp.MustCompile(`\d+\.\d+`)
	fmt.Println(reg5.FindAllStringSubmatch(str2, -1))

	/*匹配电子邮箱*/
	str3 := "zhengzhistudio@gmail.com"
	reg6 := regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	fmt.Println("电子邮箱为", reg6.FindAllStringSubmatch(str3, -1))

	/*匹配电子邮箱地址*/
	/*^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$*/
	/*\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)**/

}
