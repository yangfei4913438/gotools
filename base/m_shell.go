package base

import (
	"github.com/codeskyblue/go-sh"
)

//使用方法，有目录写目录，没有目录用"."或者""表示当前目录
func ShExec(path, cmd string, args ...string) (string, string) {
	Session := sh.NewSession()
	Session.ShowCMD = false
	Session.SetDir(path)
	out, err := Session.Command(cmd, args).Output()
	if err != nil {
		return "error", err.Error()
	} else {
		return "ok", string(out)
	}
}

/*
注意：对于下面这种情况的实现，可能无法准确判断执行是否成功，需要主动提供判断结果，例如查询sql进行匹配。
sqlplus sys/xxxx <<EOF
sql info
EOF

针对正常的shell命令操作，基本上测试下来是没有任何问题的。

还有就是返回值，这里因为涉及到输出，所以没有使用错误类型的返回。只返回了字符串判断。
*/

/*
例子1：
func main() {
	err, out := ShExec("Unzip", "sh", "test.sh")
	println(err)
	println(out)
}
输出1：
ok

----------------------------------------------------------------------------
例子2：
func main() {
	err, out := ShExec("ls", "-l")
	println(err)
	println(out)
}
输出2：
error
exec: "-l": executable file not found in $PATH

----------------------------------------------------------------------------
例子3：
func main() {
	err, out = ShExec(".", "ls", "-l")
	println(err)
	println(out)
}
输出3：
ok
total 37816
drwxr-xr-x   3 yangfei  staff       102 Dec  6 15:42 Download
-rw-r--r--   1 yangfei  staff       976 Dec  7 13:37 README.md
drwxr-xr-x   5 yangfei  staff       170 Dec  7 14:14 Unzip
-rwxr-xr-x   1 yangfei  staff  19273584 Dec  6 17:45 auto-upgrade
-rw-r--r--   1 yangfei  staff       419 Dec  6 15:18 auto-upgrade.iml
-rw-r--r--   1 yangfei  staff      4053 Dec  6 15:18 auto-upgrade.ipr
-rw-r--r--   1 yangfei  staff     67913 Dec  7 14:13 auto-upgrade.iws
drwxr-xr-x   5 yangfei  staff       170 Dec  7 11:25 conf
drwxr-xr-x  14 yangfei  staff       476 Dec  7 09:49 controllers
drwxr-xr-x   4 yangfei  staff       136 Dec  6 15:42 logs
-rw-r--r--   1 yangfei  staff       610 Dec  7 14:22 main.go
drwxr-xr-x  25 yangfei  staff       850 Dec  7 14:22 models
drwxr-xr-x   3 yangfei  staff       102 Dec  6 17:19 routers
drwxr-xr-x   4 yangfei  staff       136 Dec  6 15:42 sqls
drwxr-xr-x   3 yangfei  staff       102 Dec  6 15:42 tests
drwxr-xr-x   3 yangfei  staff       102 Dec  6 15:42 views

*/
