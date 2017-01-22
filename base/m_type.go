package base

import (
	"strconv"
)

func IntToStr(x int) (string, error) {
	return strconv.Itoa(x), nil
}

func Int64ToStr(x int64) (string, error) {
	return strconv.FormatInt(x, 10), nil
}

func StrToInt(x string) (int, error) {
	return strconv.Atoi(x)
}

func StrToInt64(x string) (int64, error) {
	return strconv.ParseInt(x, 10, 64)
}

func Float64ToStr(x float64) (string, error) {
	return strconv.FormatFloat(x, 'f', 2, 64), nil
}

func Float32ToStr(x float64) (string, error) {
	return strconv.FormatFloat(x, 'f', 2, 32), nil
}
