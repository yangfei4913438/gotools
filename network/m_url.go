package network

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

//检查网络状态,超时时间为5秒
//这里需要加上端口,且不能添加如：http:// 这样的网络协议
func UrlPing(url string) (string, error) {
	r, err := net.DialTimeout("tcp", url, time.Second*5)
	if r != nil {
		r.Close()
	}

	if err == nil {
		res := "连接url(" + url + ")成功:)"
		//网络正常,返回0
		return res, nil
	} else {
		//网络异常,返回1
		res := "连接url(" + url + ")失败: " + err.Error()
		return res, err
	}
}

//使用GET方法获取url状态, 需要完整url, 不需要端口
func UrlStatus(url string) (bool, string, error) {
	s, err := http.Get(url)
	if err != nil {
		return false, "", err
	} else {
		if s.StatusCode == 200 {
			return true, strconv.Itoa(s.StatusCode), nil
		} else {
			//只允许200，其他状态值都是异常状态
			return false, strconv.Itoa(s.StatusCode), nil
		}
	}
}
