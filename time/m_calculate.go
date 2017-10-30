package time

import (
	"github.com/yangfei4913438/gotools/base"
	"time"
)

//用于时间数字转字符串，内部函数AddMonth使用
func timeStr(t int) string {
	var tmp string
	if t < 10 {
		tmp = "0" + base.IntToStr(t)
	} else {
		tmp = base.IntToStr(t)
	}
	return tmp
}

//传入Unix时间戳，以及需要增加的月份数，返回一个新的时间对象
func AddMonth(basetime int64, md int) time.Time {

	bt := time.Unix(basetime, 0)
	y := bt.Year()
	m := int(bt.Month())
	d := bt.Day()
	hh := bt.Hour()
	mm := bt.Minute()
	ss := bt.Second()

	ny := y + md/12

	nm := m + md%12
	if nm > 12 {
		ny += 1
		nm = nm % 12
	}

	rd, _, _ := OneMonth(ny, nm)
	if d > rd {
		d = d - rd
		nm += 1
		if nm > 12 {
			ny += 1
			nm = nm % 12
		}
	}

	tmp := timeStr(ny) + "-" + timeStr(nm) + "-" + timeStr(d) + " " + timeStr(hh) + ":" + timeStr(mm) + ":" + timeStr(ss)

	new_time, _ := time.Parse("2006-01-02 15:04:05", tmp)

	return new_time
}
