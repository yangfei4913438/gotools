package base

import (
	"github.com/yangfei4913438/gotools/math"
)

func GetToken(length int, ismd5 bool) string {

	res := ""

	for i := 0; i < length; i++ {
		res += math.RandStr()
	}

	if ismd5 {
		return StrMD5(res)
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
