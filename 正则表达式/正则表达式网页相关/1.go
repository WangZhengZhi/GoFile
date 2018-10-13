package main

import (
	"fmt"
	"regexp"
)

/* 			``	这是原生字符串
 */
func main() {
	str := `

<html>
<head>
    <body>


        <h1>
            这是一段示范用的例子
            有中文
            有english
           	有123456789
        </h1>
    </body>
</head>

</html>
    






`
	reg := regexp.MustCompile(`[\p{Han}]+`)
	fmt.Println("中文有：", reg.FindAllStringSubmatch(str, -1))
	reg1 := regexp.MustCompile(`有.+`)
	fmt.Println("匹配`有XXXXXXX`", reg1.FindAllStringSubmatch(str, -1))
	reg2 := regexp.MustCompile(`[\d]`)
	fmt.Println("数字为：", reg2.FindAllString(str, -1))
	reg3 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`) /*只要<h1></h1>中间的*/
	/*?s:是换行也要*/
	/*.*?重复匹配前面的字段，越少越好(优先退出重复)*/
	//fmt.Println("<h1>between</h1>:", reg3.FindAllString(str, -1))
	/*过滤<> </>*/
	result := reg3.FindAllStringSubmatch(str, -1)
	for _, data := range result {

		fmt.Println(data[1]) /*不带<></>*/
		fmt.Println(data[0]) /*带<></>*/
	}

}
