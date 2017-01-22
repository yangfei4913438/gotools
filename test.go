package main

import (
	"fmt"
	"gotools/base"
)

func main() {
	fmt.Println("this is shell test, begin!\n")
	sh_res, sh_out := base.ShExec(".", "pwd")
	fmt.Println("res: ", sh_res)
	fmt.Println("out: ", sh_out)
	fmt.Println("this is shell test, end!\n")

	str_res, str_err := base.StrToInt64("222")
	if str_err != nil {
		fmt.Println(str_err.Error() + "\n")
	} else {
		fmt.Println(str_res, "\n")
	}

	fmt.Println(base.ErrorCustom("自定义错误信息").Error())

}
