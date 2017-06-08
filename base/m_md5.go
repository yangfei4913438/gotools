package base

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func StrMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func StrSHA256(str string) string {
	hash_value := sha256.New()
	hash_value.Write([]byte(str))
	md := hash_value.Sum(nil)
	return hex.EncodeToString(md)
}
