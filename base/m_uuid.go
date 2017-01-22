package base

import (
	"fmt"
	"net"
	"github.com/nu7hatch/gouuid"
)

//获取机器码
func GetUUID(org_id, machine_type, regist_number string) (string, int) {
	_, mac_addr := AllMac()
	str := org_id + "-" + machine_type + "-" + regist_number + "-" + mac_addr
	u5, err := uuid.NewV5(uuid.NamespaceURL, []byte(str))
	if err != nil {
		fmt.Println("获取UUID失败:" + err.Error())
		return "", 1
	}
	return u5.String() + "-" + org_id + "-" + machine_type, 0
}

func AllMac() (string, string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "error", err.Error()
	}
	var mac_all_addr string
	for _, inter := range interfaces {
		if len(inter.HardwareAddr.String()) > 0 {
			wk := string([]byte(inter.Name)[len(inter.Name)-1])
			if wk == "0" {
				mac_all_addr = mac_all_addr + fmt.Sprint(inter.Name+":"+inter.HardwareAddr.String()+"-")
			} else if wk == "1" {
				mac_all_addr = mac_all_addr + fmt.Sprint(inter.Name+":"+inter.HardwareAddr.String()+"-")
			} else if wk == "2" {
				mac_all_addr = mac_all_addr + fmt.Sprint(inter.Name+":"+inter.HardwareAddr.String()+"-")
			} else if wk == "3" {
				mac_all_addr = mac_all_addr + fmt.Sprint(inter.Name+":"+inter.HardwareAddr.String()+"-")
			}
		}
	}
	if len(mac_all_addr) == 0 {
		//没有匹配的网卡，就返回一组指定的字符串
		return "ok", "00-0000-0000-0000-0000-0000"
	} else {
		return "ok", string([]byte(mac_all_addr)[:len(mac_all_addr)-1])
	}
}
