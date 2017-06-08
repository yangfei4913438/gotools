# Golang Tools [![wercker status](https://app.wercker.com/status/5fcc7df5c1d51417a43c381aa0ec3de6/s/master "wercker status")](https://app.wercker.com/project/byKey/5fcc7df5c1d51417a43c381aa0ec3de6)

The role of the project: as the basis of other golang project module.

## How to install
Use `go get` to install or upgrade (`-u`) the `gotools` package.

    go get -u github.com/yangfei4913438/gotools

## Usage
Base Example: 

```go
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

	//求哈希值
	fmt.Println(base.StrMD5("hello"))
	fmt.Println(base.StrSHA1("hello"))
	fmt.Println(base.StrSHA256("hello"))
	fmt.Println(base.StrSHA512("hello"))
}
```

Info Example: 
    
    Please be patient...
