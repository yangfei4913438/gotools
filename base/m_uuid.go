package base

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"net"
)

//获取机器码
func GetUUID(org_id, machine_type, regist_number string) (string, error) {
	mac_addr, mac_err := AllMac()
	if mac_err != nil {
		return "", mac_err
	}
	str := org_id + "-" + machine_type + "-" + regist_number + "-" + mac_addr
	u5, err := uuid.NewV5(uuid.NamespaceURL, []byte(str))
	if err != nil {
		return "", err
	}
	return u5.String() + "-" + org_id + "-" + machine_type, nil
}

func AllMac() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
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
		//没有匹配的网卡，就返回错误信息
		return "", ErrorCustom("机器中不存在网卡!")
	} else {
		return string([]byte(mac_all_addr)[:len(mac_all_addr)-1]), nil
	}
}
