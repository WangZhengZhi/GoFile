package main

import (
	"encoding/json"
	"fmt"
)

/*json解析到结构体*/
/*解码json*/
/*需要二次编码*/
type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	Isok     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	jsondata := `

{
 	"company": "Itcast",
 	"isok": true,
 	"price": 99999.9999,
 	"subjects": [
 		"Go",
 		"java",
 		"cpp",
 		"python",
 		"Test"
 	]
 }

`
	var tmp IT                                    //定义一个结构体变量
	err := json.Unmarshal([]byte(jsondata), &tmp) //记得是地址传递&
	if err != nil {
		fmt.Println("err:=", err)
	}
	fmt.Printf("tmp=%+v\n", tmp)
	/*也可以只解析一个*/
	type IT2 struct {
		Company string `json:company` /*只要一个数据*/
	}
	var tmp2 IT2
	err = json.Unmarshal([]byte(jsondata), &tmp2)
	if err != nil {
		fmt.Println("err:=", err)
		return
	}
	fmt.Printf("%+v\n", tmp2)

}
