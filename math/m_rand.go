package math

import (
	"math/rand"
	"time"
)

//生成0-9之间的随机数
func RandInt() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10)
}
