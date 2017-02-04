package time

import "time"

//完整的日期和时间
func GetFullTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//获取日期，不带时间
func GetData() string {
	return time.Now().Format("2006-01-02")
}

//获取时间，不带日期
func GetTime() string {
	return time.Now().Format("15:04:05")
}

/*
例子：
fmt.Println(time.GetFullTime())  // 2017-02-04 18:45:53
fmt.Println(time.GetData())      // 2017-02-04
fmt.Println(time.GetTime())      // 18:45:53
*/
