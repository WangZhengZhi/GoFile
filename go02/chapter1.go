package main

//switch 的使用
import (
	"fmt"
)

func main() {
	var a int
	a = 10
	switch a {
	case 1:
		fmt.Println("这里是1")
	case 2:
		fmt.Println("这里是2")
	default:
		fmt.Println("不满足所有case") //switch中不满足所有的case 便会只调用default 情况
		//使用switch 可以使用逗号分隔的列表
		var word byte
		word = '+'
		switch word {
		case '*', '/', '-':
			fmt.Println("不满足 - / *") //case可以使用逗号分隔开判断条件,可以使用多个
		default:
			fmt.Println("满足+")
		}
	}

}

//比较两个字节数组字典数序先后的整数
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}

//a==b->0
//a>b->+1
//a<b->-1
