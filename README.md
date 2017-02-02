# Golang Tools [![wercker status](https://app.wercker.com/status/49eb7e5141480529ed7d4e53fe13ab00/s/master "wercker status")](https://app.wercker.com/project/byKey/49eb7e5141480529ed7d4e53fe13ab00)

The role of the project: as the basis of other golang project module.

## How to install
    go get github.com/yangfei4913438/gotools

## Usage
Base Example: 

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
    
    	//这里写的download相对路径，需要创建真实目录
    	network.UrlDownload("download", "http://dldir1.qq.com/qqfile/QQforMac/QQ_V5.4.1.dmg")
    
    }

Info Example: 
    
    Please be patient...
