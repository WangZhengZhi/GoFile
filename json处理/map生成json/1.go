package main

import (
	"encoding/json"
	"fmt"
)

/*通过map生成json*/
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
func main() {
	//创建一个MAP
	m := make(map[string]interface{}, 4) /*interface{}万能接口，可以接受任何数据*/
	m["company"] = "Itcast"
	m["subjects"] = []string{"Go", "java", "cpp", "python", "Test"}
	m["isok"] = true
	m["price"] = 99999.9999
	/*编码成json*/
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("err:=", err)
		return
	}
	fmt.Println(string(data))
	/*输出格式化的形式*/
	data, err = json.MarshalIndent(m, " ", "	")
	if err != nil {
		fmt.Println("err:=", err)
		return
	}
	fmt.Println(string(data))

}
