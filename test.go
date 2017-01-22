package main

import (
	"fmt"
	"gotools/base"
)

func main() {
	sh_out, sh_err := base.ShExec("", "ls", "-l")
	if sh_err != nil {
		fmt.Println("error: ", sh_err.Error())
	} else {
		fmt.Println("out: ", sh_out)
	}

	fmt.Println(base.ErrorCustom("自定义错误信息").Error())

}
