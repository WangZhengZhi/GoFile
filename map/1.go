package main

import "fmt"

//map的声明方式
//var map_varible map[key_data_type]value_type
//使用make的方式
//map_varible:=make([key_data_type]value_type)
//如果不初始化，这些map是空的nil
func main() {
	maps := make(map[string]string)
	for key, value := range maps {
		fmt.Println(key, value)
	}                   //空的map不能存放数据,nil
	maps["1"] = "hello" //给map增加数据
	delete(maps, "1")
	for key, value := range maps {
		fmt.Println(key, value)
		fmt.Println("-------")
	}
	maps["河南"] = "郑州"
	maps["河北"] = "石家庄"
	maps["山东"] = "临沂"
	maps["广东"] = "深圳"
	maps["usa"] = "DC"
	for key, value := range maps {
		fmt.Println(key, value)
	} //给maps赋值
	/*删除maps中的数据*/
	delete(maps, "usa")
	printMap(maps) //调用打印函数
	/*查看map中是否有某个元素*/
	key, ok := maps["河南"]
	if ok {
		fmt.Println(key)
	} else {
		fmt.Println("不存在")
	}

}
func printMap(maps map[string]string) {
	fmt.Println("开始打印......")
	for key, value := range maps {
		fmt.Println(key, value)

	}
	fmt.Println("结束打印...")

} //打印map的函数
