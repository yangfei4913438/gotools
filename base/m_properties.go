package base

import "strings"

//获取properties文件中KEY的值
func Properties_Get_Value(file_path, key string) string {
	var return_info string
	info, res := Readline(file_path)
	if len(info) == 0 {
		one_info, _ := ReadAll(file_path)
		str_list := strings.Split(string(one_info), "=")
		if CleanSpace(str_list[0]) == key {
			return_info = CleanSpace(str_list[1])
		}
	} else {
		if res == 0 {
			for _, i := range info {
				str_list := strings.Split(i, "=")
				if CleanSpace(str_list[0]) == key {
					return_info = CleanSpace(str_list[1])
				}
			}
		}
	}
	return strings.Replace(return_info, "\n", "", -1)
}