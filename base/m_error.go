package base

import "errors"

//自定义错误信息
func ErrorCustom(text string) error {
	return errors.New(text)
}

//使用说明
// fmt.Println(ErrorCustom("自定义错误信息").Error())
