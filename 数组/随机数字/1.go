package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机数字的生成
func main() {
	//设置种子
	//如果种子参数一样的，每次运行的结果也是一样的
	rand.Seed(time.Now().UnixNano()) //使用系统时间作为种子参数
	for i := 0; i < 5; i++ {
		//产生随机数字
		fmt.Println("rand=", rand.Int())     //生成很大的整数
		fmt.Println("rand=", rand.Intn(100)) //限制在100以内
	}
}
