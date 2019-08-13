package validator

import (
	"github.com/astaxie/beego"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type validator struct{}

var Validate validator

// 验证是否为空
func (v *validator) CheckBlank(val interface{}) bool {
	value := reflect.ValueOf(val)
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// 验证是否为时间字符串
func (v *validator) CheckTimeString(val string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", val)
	if err != nil {
		beego.Error("校验时间字符串出错:", err)
		return false
	}
	return true
}

// 验证时区字符串
func (v *validator) CheckZone(zone int) bool {
	if zone < -12 || zone > 12 {
		beego.Error("错误的时区值:", zone)
		return false
	}
	return true
}

// 验证用户名的长度
func (v *validator) CheckUserName(name string) bool {
	// 用户名的长度必须大于等于 2个字符，且小于等于 20个字符。
	res := len(name) >= 2 && len(name) <= 20
	if res {
		return true
	} else {
		beego.Error("用户名的长度必须大于等于 2个字符，且小于等于 20个字符！")
		return false
	}
}

// 验证密码
func (v *validator) CheckPassword(pwd string) bool {
	// 规则 1：必须含有数字 0-9
	reg1 := regexp.MustCompile(`[0-9]+`)

	// 规则 2：必须含有大写字母 A-Z
	reg2 := regexp.MustCompile(`[A-Z]+`)

	// 规则 3：必须含有小写字母 a-z
	reg3 := regexp.MustCompile(`[a-z]+`)

	// 规则 4：必须含有指定的特殊字符
	reg4 := regexp.MustCompile(`[-=[;,./~!@#$%^*()_+}{:?]+`)

	// 规则 5：密码长度必须在 6-20 位之间
	reg5 := regexp.MustCompile(`^[\s\S]{6,20}$`)

	// 必须同时满足 5 种规则
	res := reg1.MatchString(pwd) && reg2.MatchString(pwd) && reg3.MatchString(pwd) && reg4.MatchString(pwd) && reg5.MatchString(pwd)

	if res {
		return true
	} else {
		beego.Error("密码必须含有至少: 一个数字、一个小写字母、一个大写字母、一个特殊符号，并且长度在 6-20 位之间！")
		return false
	}
}

// 验证电子邮件
func (v *validator) CheckEmail(email string) bool {
	if email == "" {
		// 没有填写 email 的情况下，判断通过，因为允许为空
		return true
	}
	reg := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	res := reg.MatchString(email)
	if res {
		return true
	} else {
		beego.Error("电子邮件的格式错误!")
		return false
	}
}

// 验证 IP 地址
func (v *validator) CheckIP(ip string) bool {
	if ip == "" {
		beego.Error("IP地址不允许为空!")
		return false
	}
	ipList := strings.Split(ip, ".")
	ip1, err := strconv.Atoi(ipList[0])
	if err != nil {
		beego.Error("错误的 IP 格式:", err)
		return false
	}
	ip2, err := strconv.Atoi(ipList[1])
	if err != nil {
		beego.Error("错误的 IP 格式:", err)
		return false
	}
	ip3, err := strconv.Atoi(ipList[2])
	if err != nil {
		beego.Error("错误的 IP 格式:", err)
		return false
	}
	ip4, err := strconv.Atoi(ipList[3])
	if err != nil {
		beego.Error("错误的 IP 格式:", err)
		return false
	}
	// 每一个值的取值范围，都必须在 0-255 之间
	if ip1 < 0 || ip1 > 255 || ip2 < 0 || ip2 > 255 || ip3 < 0 || ip3 > 255 || ip4 < 0 || ip4 > 255 {
		beego.Error("错误的 IP 格式:", ip)
		return false
	}
	// 0、127 开头的 IP, 以及大于等于 224 开头的 IP，都是保留地址
	if ip1 == 0 || ip1 == 127 || ip1 >= 224 {
		// 0.0.0.0 是允许使用的默认 IP
		if ip != "0.0.0.0" {
			beego.Error("保留的 IP 地址:", ip)
			return false
		}
	}
	return true
}
