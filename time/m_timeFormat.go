package time

import (
	"fmt"
	"strconv"
	"time"
)

//参数：时间戳，时区(-12到12)。获取相应时区的时间
func GetTimeZoneTime(timestamp int64, value int) time.Time {
	//需要转换成UTC时间，就是0时区的时间
	return time.Unix(timestamp, 0).UTC().Add(time.Duration(value) * time.Hour)
}

type FormatTime struct {
	time.Time
}

func (t *FormatTime) ToString() string {
	return t.Format("2006-01-02 15:04:05")
}

func (t *FormatTime) ToUnix() int64 {
	//英文时间是东八区的时间，所以去掉8个小时，再转unix就对了
	return t.UTC().Add(-8 * OneHour).Unix()
}

//字符串格式化为时间
func FormatStrTime(time_str string) (*FormatTime, error) {
	x, err := time.Parse("2006-01-02 15:04:05", time_str)
	if err != nil {
		return nil, err
	}
	return &FormatTime{x}, nil
}

//时间戳格式化为字符串，这里默认为东八区的时间戳
func FormatUnixTime(t int64) string {
	//默认是东八区的时间戳
	return time.Unix(t, 0).UTC().Add(time.Hour * 8).Format("2006-01-02 15:04:05")
}

//获取时区显示字符串
func GetTimeZone(value int) string {
	if value >= 0 {
		return "(GMT +0" + strconv.Itoa(value) + ":00)"
	} else {
		return "(GMT -0" + strconv.Itoa(value) + ":00)"
	}
}

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

//获取日期相对的周几
func GetWeekDay(Year, Month, Day string) string {
	if len(Month) == 1 {
		Month = "0" + Month
	}
	if len(Day) == 1 {
		Day = "0" + Day
	}

	DATE := fmt.Sprintf("%s-%s-%s", Year, Month, Day)
	t, _ := time.Parse("2006-01-02", DATE)
	return t.Weekday().String()
}

/*
例子：
fmt.Println(time.GetFullTime())  // 2017-02-04 18:45:53
fmt.Println(time.GetData())      // 2017-02-04
fmt.Println(time.GetTime())      // 18:45:53
fmt.Println(GetWeekDay("2017", "3","1"))  //Wednesday
*/
