package main

import (
	"encoding/json"
	"fmt"
)

/*json源数据*/
/*
"company":"itcast",
"subject":[
"GO"
"CPP"
"Java"
"Test"
],
"isok":true,
"price":8888888.999





*/

/*通过结构体生成json*/
/*成员变量名必须是大写*/
type IT struct {
	Company string `json:"company"` /*输出的json格式为小写,这是二次编码*/
	Subject []string
	Isok    bool    `json:",string"` /*输出为string类型*/
	Price   float64 `json:"-"`       /*字段不会输出到屏幕*/
}

func main() {
	/*定义一个结构体变量，同时初始化*/
	s := IT{"Itcast", []string{"go", "c++", "python", "test"}, true, 8888}
	/*编码，根据内容生成json*/
	str, err := json.Marshal(s) /*str是切片要进行转化*/
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("json=", string(str))
	/*格式化编码*/
	str, err = json.MarshalIndent(s, " ", "		") /*str是切片要进行转化*/
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("json=", string(str)) /*输出有格式的感觉*/
}
