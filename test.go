package main

import (
	"fmt"
	"gotools/base"
	"gotools/network"
)

func main() {
	sh_out, sh_err := base.ShExec("", "ls", "-l")
	if sh_err != nil {
		fmt.Println("error: ", sh_err.Error())
	} else {
		fmt.Println("out: ", sh_out)
	}

	fmt.Println(base.ErrorCustom("自定义错误信息").Error())

	//这里写的download相对路径，需要创建真实目录
	network.UrlDownload("download", "http://dldir1.qq.com/qqfile/QQforMac/QQ_V5.4.1.dmg")

}
