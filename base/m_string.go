package base

import (
	"strings"
)

//截取字符串，参数：字符串，开始下标，结束下标[包含结束下标对应的值]
func Splitstr(str string, start int, end int) string {
	end_num := end + 1
	if end_num > len(str) {
		end_num = len(str)
	}
	return string([]byte(str)[start:end_num])
}

//清除字符串中的空白
func CleanSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}

//检查字符串中是否包含某个字符,支持单字符和词汇的检查
func StrContains(str, obj string) bool {
	return strings.Contains(str, obj)
}

//检查字符或字符串出现的次数
func StrCount(str, obj string) int {
	return strings.Count(str, obj)
}
