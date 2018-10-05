package main

import "fmt"

/*
go语言中的错误处理
error是一个接口类型
定义如下：
type error interface{
Error()string
}
*/
func main() {
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10:", result)
	} //正常的情况
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is:", errorMsg)
	} //除数是0的情况下
}

/*func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.NewErrCopyResource("math:square root of negative number")
	}
}*/
type DivideError struct {
	dividee int
	divider int
}

//实现error接口
func (de *DivideError) Error() string {
	strFormat := `can not proceed,the divider is zero
dividee:%d
divier:0`
	return fmt.Sprintf(strFormat, de.dividee)
} //使用指针接口才可以改变数值

//定义‘int’类型得的除法运算操作函数
func Divide(varDividee int, varDivider int) (result int, errMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}

}
