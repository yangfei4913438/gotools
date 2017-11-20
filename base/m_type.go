package base

import (
	"strconv"
)

func IntToStr(x int) string {
	return strconv.Itoa(x)
}

func Int64ToStr(x int64) string {
	return strconv.FormatInt(x, 10)
}

func StrToInt(x string) (int, error) {
	return strconv.Atoi(x)
}

func StrToInt64(x string) (int64, error) {
	return strconv.ParseInt(x, 10, 64)
}

func Float64ToStr(x float64) string {
	return strconv.FormatFloat(x, 'f', 2, 64)
}

func StrToFloat64(x string) (float64, error) {
	return strconv.ParseFloat(x, 64)
}
