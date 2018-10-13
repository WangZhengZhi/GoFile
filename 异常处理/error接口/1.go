package main

/*error问题
普通错误*/
import (
	"fmt"
	"github.com/pkg/errors"
)

/*error接口示例*/
func Mydiv(numa, numb int) (result int, err error) {
	err = nil
	if numb == 0 {
		err = errors.New("除数不能为0")
	} else {
		result = numa / numb
	}
	return
}
func main() {

	result, err := Mydiv(10, 0)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("result:", result)
	}
}
