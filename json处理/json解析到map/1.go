package main

import (
	"encoding/json"
	"fmt"
)

/*json解析到map*/

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
	m := make(map[string]interface{}, 4)        /*map 创建*/
	err := json.Unmarshal([]byte(jsondata), &m) /*记得是地址传递*/
	if err != nil {
		fmt.Println("err:=", err)
		return
	}
	fmt.Printf("%+v:\n", m)
	/*类型断言取出元素*/
	for key, value := range m {
		//fmt.Printf("%v------%v\n", key, value)
		switch data := value.(type) {
		case string:
			fmt.Printf("%v---%v\n", key, data)
		case bool:
			fmt.Printf("%v---%v\n", key, data)
		case float64:
			fmt.Printf("%v---%v\n", key, data)
		case []interface{}:
			fmt.Printf("%v---%v\n", key, data)
		}
	}

}
