package base

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func GetToken(length int, ismd5 bool) string {

	src := [...]string{
		"Q", "@", "8", "y", "%", "^", "5", "Z", "(", "G", "_", "O", "*",
		"S", "-", "N", "<", "D", "{", "}", "[", "]", "h", ";", "W", ".",
		"/", "|", ":", "1", "E", "L", "4", "&", "6", "7", "#", "9", "a",
		"A", "b", "B", "~", "C", "d", ">", "e", "2", "f", "P", "g", ")",
		"?", "H", "i", "X", "U", "J", "k", "r", "l", "3", "t", "M", "n",
		"=", "o", "+", "p", "F", "q", "!", "K", "R", "s", "c", "m", "T",
		"v", "j", "u", "V", "w", ",", "x", "I", "$", "Y", "z"}

	res := ""

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		n := r.Intn(len(src)) //取随机数
		res += src[n]
	}

	if ismd5 {
		h := md5.New()
		h.Write([]byte(res))
		return hex.EncodeToString(h.Sum(nil))
	} else {
		return res
	}

}

/*
第一个参数，是取原始值的长度；
第二个参数，是告诉代码，是否需要MD5进行加密；

范例：
//获取token,原生
fmt.Println(base.GetToken(32, false))

//获取token,md5加密的值
fmt.Println(base.GetToken(32, true))

*/