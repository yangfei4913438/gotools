package base

import (
	"github.com/codeskyblue/go-sh"
)

//使用方法，有目录写目录，没有目录用"."或者""表示当前目录
func ShExec(path, cmd string, args ...string) (string, error) {
	Session := sh.NewSession()
	Session.ShowCMD = false
	Session.SetDir(path)
	out, err := Session.Command(cmd, args).Output()
	if err != nil {
		return "", err
	} else {
		return string(out), nil
	}
}

/*
注意：对于下面这种情况的实现，可能无法准确判断执行是否成功，需要主动提供判断结果，例如查询sql进行匹配。
sqlplus sys/xxxx <<EOF
sql info
EOF

针对正常的shell命令操作，基本上测试下来是没有任何问题的。
*/

/*
例子1：
func main() {
	sh_out, sh_err := base.ShExec("base", "pwd")
	if sh_err != nil {
		fmt.Println("error: ", sh_err.Error())
	} else {
		fmt.Println("out: ", sh_out)
	}
}
输出1：
out:  /Users/yangfei/goprojects/src/gotools/base

----------------------------------------------------------------------------
例子2：
func main() {
	sh_out, sh_err := base.ShExec("base1", "pwd")
	if sh_err != nil {
		fmt.Println("error: ", sh_err.Error())
	} else {
		fmt.Println("out: ", sh_out)
	}
}
输出2：
error:  chdir base1: no such file or directory

----------------------------------------------------------------------------
例子3：
func main() {
	sh_out, sh_err := base.ShExec("", "ls", "-l")
	if sh_err != nil {
		fmt.Println("error: ", sh_err.Error())
	} else {
		fmt.Println("out: ", sh_out)
	}
}
输出3：
out:  total 232
-rw-r--r--   1 yangfei  staff   7651 Jan 22 07:19 LICENSE
-rw-r--r--   1 yangfei  staff    186 Jan 22 08:31 README.md
drwxr-xr-x  11 yangfei  staff    374 Jan 23 06:34 base
-rw-r--r--   1 yangfei  staff    806 Jan 22 08:07 gotools.iml
-rw-r--r--   1 yangfei  staff  36022 Jan 23 06:02 gotools.ipr
-rw-r--r--   1 yangfei  staff  59386 Jan 23 06:21 gotools.iws
drwxr-xr-x   5 yangfei  staff    170 Jan 23 05:21 network
-rw-r--r--   1 yangfei  staff    283 Jan 23 06:34 test.go
drwxr-xr-x   6 yangfei  staff    204 Jan 23 05:21 time

*/
