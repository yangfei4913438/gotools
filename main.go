package main

import (
	"fmt"
	"github.com/yangfei4913438/gotools/network"
)

func main() {

	//这里写的download相对路径，需要创建真实目录
	err := network.UrlDownload("download", "http://dldir1.qq.com/qqfile/QQforMac/QQ_V5.4.1.dmg")
	if err != nil {
		fmt.Println(err)
	}
}
