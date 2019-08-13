package database

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/yangfei4913438/gotools/base"
	"github.com/yangfei4913438/gotools/time"
	"golang.org/x/crypto/scrypt"
)

// 给密码加密
func PasswordEncryption(username string, password string, createAt int64) (string, error) {
	// 生成盐[用户名、密码、用户创建时间(时间戳类型)]
	salt := base.StrMD5(username + password + time.TimestampToLocal(createAt) + "v1.0.0")

	// 获取加密结果
	dk, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	if err != nil {
		beego.Error("密码加密出错:", err)
		return "", err
	}

	// 返回加密数据, 因为上面得到的是一个哈希类型的数据，所以要进行格式化处理
	return fmt.Sprintf("%x", dk), nil
}
