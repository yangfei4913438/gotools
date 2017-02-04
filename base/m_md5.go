package base

import (
	"crypto/md5"
	"encoding/hex"
)

func StrMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}