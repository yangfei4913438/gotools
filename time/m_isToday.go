package time

import "time"

func IsToday(newtime string) bool {
	//注意：这里的格式非常苛刻，日期单数只能写2017-05-06，绝对不能写成2017-5-6，这种参数绝对是错误的！

	new_time, _ := time.Parse("2006-01-02", newtime)
	today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	if new_time.Unix() == today.Unix() {
		return true
	} else {
		return false
	}
}
