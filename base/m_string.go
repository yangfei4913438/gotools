package base

import "strings"

//截取字符串，参数：字符串，开始下标，结束下标[包含结束下标对应的值]
func Splitstr(str string, start int, end int) string {
	end_num := end + 1
	return string([]byte(str)[start:end_num])
}

//清除字符串中的空白
func CleanSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}
