package network

import (
	"net"
	"net/http"
	"strconv"
	"time"
)

//检查网络状态,超时时间为3秒
func UrlPing(url string) (string, error) {
	_, err := net.DialTimeout("tcp", url, time.Second*5)
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

//使用GET方法获取url状态(限ip,对域名无效)
//1表示异常，0表示正常。用于正常状态下的数值判断
func UrlStatus(url string) (int, string, error) {
	s, err := http.Get(url)
	if err != nil {
		return 1, "", err
	} else {
		if s.StatusCode == 200 {
			return 0, strconv.Itoa(s.StatusCode), nil
		} else {
			//只允许200，其他状态值都是异常状态
			return 1, strconv.Itoa(s.StatusCode), nil
		}
	}
}
