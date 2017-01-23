package base

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"strings"
)

//截取字符串，参数：字符串，开始下标，结束下标[包含结束下标对应的值]
func Splitstr(str string, start int, end int) string {
	end_num := end + 1
	return string([]byte(str)[start:end_num])
}

//清除字符串中的空白
func CleanSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}

//检查字符串中是否包含某个字符
//不支持多个字符的包含检查
func StrContains(str, obj string) bool {
	var a []interface{}
	for _, v := range str {
		a = append(a, fmt.Sprintf("%c", v))
	}
	class := mapset.NewSetFromSlice(a)
	return class.Contains(obj)
}
