package time

import "time"

func TimeDiff(oldTime, newTime string) int {
	//传入的时间，必须是规定的格式 "2006-01-02 15:04:05"
	//用于计算2个时间之间有多少秒的差别

	old_time, _ := time.Parse("2006-01-02 15:04:05", oldTime)
	new_time, _ := time.Parse("2006-01-02 15:04:05", newTime)

	if old_time.Unix() >= new_time.Unix() {
		subTimes := old_time.Sub(new_time).Seconds()
		return int(subTimes)
	} else {
		subTimes := new_time.Sub(old_time).Seconds()
		return int(subTimes)
	}
}

//参数：时间戳，时区(-12到12)。获取相应时区的时间
func GetTimeZoneTime(timestamp int64 ,value int) time.Time {
	//需要转换成UTC时间，就是0时区的时间
	return time.Unix(timestamp,0).UTC().Add(time.Duration(value)*time.Hour)
}
