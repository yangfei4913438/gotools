package math

import (
	"github.com/yangfei4913438/gotools/base"
)

//单位转换: 传入值的单位(B)
func MathChangeUnit(unit int64) string {
	m := int64(1024)
	switch {
	case unit < m:
		return base.Int64ToStr(unit) + "B"
	case m <= unit && unit < m*m:
		tmp := float64(unit) / 1024
		return base.Float64ToStr(tmp) + "KB"
	case m*m <= unit && unit < m*m*m:
		tmp := float64(unit) / 1024 / 1024
		return base.Float64ToStr(tmp) + "MB"
	case m*m*m <= unit && unit < m*m*m*m:
		tmp := float64(unit) / 1024 / 1024 / 1024
		return base.Float64ToStr(tmp) + "GB"
	default:
		tmp := float64(unit) / 1024 / 1024 / 1024 / 1024
		return base.Float64ToStr(tmp) + "TB"
	}
}
