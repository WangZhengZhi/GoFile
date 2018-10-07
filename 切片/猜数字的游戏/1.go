package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*产生一个四位数字，作为猜数游戏*/
func randomInit(p *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		} //限制随机数字为四位数字
	}
	*p = num
}

/*
初始化随机数字
 */

func getNum(slice []int, num int) {
	slice[0] = num / 1000       //千位
	slice[1] = num % 1000 / 100 //百位
	slice[2] = num % 100 / 10   //十位
	slice[3] = num % 10         //个位

}

//将数字拆成切片的4个元素

func Ongame(randSlice []int, randNum int) {
	var num int
	keySlice := make([]int, 4)
	for {
		for {
			fmt.Println("请输入一个四位数")
			fmt.Scan(&num)
			if num > 999 && num < 10000 {

				break
			}
			fmt.Println("输入数字不符合要求")
		}

		getNum(keySlice, num) //将数字拆成slice的4个元素
		n := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > randSlice[i] {
				fmt.Printf("第%d位大了\n", i+1)
			} else if keySlice[i] < randSlice[i] {
				fmt.Printf("第%d位小了\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				n++
			}

		}

		if n == 4 {
			fmt.Println("全部猜对了")
			break //跳出循环
		}
	}

}

func main() {
	var randNum int
	randomInit(&randNum)
	fmt.Println("random num is:", randNum)
	randSlice := make([]int, 4)
	getNum(randSlice, randNum)
	fmt.Println(randSlice)
	Ongame(randSlice, randNum)

}
