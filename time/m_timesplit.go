package time

//第一个返回值代表分，第二个返回值代表秒
func GetMinute(seconds int) (int, int) {
	if seconds < 60 {
		return 0, seconds
	} else {
		m := seconds / 60
		s := seconds % 60
		return m, s
	}
}

//第一个返回小时，第二个分钟，第三个秒
func GetHour(seconds int) (int, int, int) {
	m, s := GetMinute(seconds)
	if m < 60 {
		return 0, m, s
	} else {
		h := m / 60
		m := m % 60
		return h, m, s
	}
}

//第一个返回天数，第二个小时，第三个分钟，第四个秒
func GetDay(seconds int) (int, int, int, int) {
	h, m, s := GetHour(seconds)
	if h < 24 {
		return 0, h, m, s
	} else {
		d := h / 24
		h := h % 24
		return d, h, m, s
	}
}

//第一个返回周，第二个天，第三个小时，第四个分钟，第五个秒
func GetWeek(seconds int) (int, int, int, int, int) {
	d, h, m, s := GetDay(seconds)
	if d < 7 {
		return 0, d, h, m, s
	} else {
		w := d / 7
		d := d % 7
		return w, d, h, m, s
	}
}
