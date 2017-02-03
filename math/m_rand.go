package math

import (
	"math/rand"
	"time"
)

//生成指定范围内的随机数【0到x】
func RandInt(x int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(x)
}

//生成0-9之间的随机数
func Rand09() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10)
}

//生成随机字符
func RandStr() string {
	src := [...]string{
		"Q", "@", "8", "y", "%", "^", "5", "Z", "(", "G", "_", "O", "*",
		"S", "-", "N", "<", "D", "{", "}", "[", "]", "h", ";", "W", ".",
		"/", "|", ":", "1", "E", "L", "4", "&", "6", "7", "#", "9", "a",
		"A", "b", "B", "~", "C", "d", ">", "e", "2", "f", "P", "g", ")",
		"?", "H", "i", "X", "U", "J", "k", "r", "l", "3", "t", "M", "n",
		"=", "o", "+", "p", "F", "q", "!", "K", "R", "s", "c", "m", "T",
		"v", "j", "u", "V", "w", ",", "x", "I", "$", "Y", "z"}

	return src[RandInt(len(src))]
}
