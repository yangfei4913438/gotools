package main

import (
	"fmt"
	"github.com/yangfei4913438/gotools/base"
	"github.com/yangfei4913438/gotools/network"
)

func main() {
	sh_out, sh_err := base.ShExec("", "ls", "-l")
	if sh_err != nil {
		fmt.Println("error: ", sh_err.Error())
	} else {
		fmt.Println("out: ", sh_out)
	}

	fmt.Println(base.ErrorCustom("自定义错误信息").Error())

	//获取token,默认32个随机值
	fmt.Println(base.GetToken(32, false))

	//获取token,将32个随机值用MD5进行加密
	fmt.Println(base.GetToken(32, true))

	//这里写的download相对路径，需要创建真实目录
	network.UrlDownload("download", "http://dldir1.qq.com/qqfile/QQforMac/QQ_V5.4.1.dmg")

}
