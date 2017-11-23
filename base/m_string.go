package base

import (
	"strings"
)

//截取字符串，参数：字符串，开始顺序[第一个数开始算，不是索引，索引是从0开始，顺序是从1开始]，结束顺序[第几个字符，包含这个字符]
func Splitstr(str string, start int, end int) string {
	start--
	end--
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

//检查字符串是不是以某个字符或字符串开始的
func StrHasPrefix(str, prefix string) bool {
	return strings.HasPrefix(str, prefix)
}

//检查字符串是不是以某个字符或字符串结尾的
func StrHasSuffix(str, prefix string) bool {
	return strings.HasSuffix(str, prefix)
}
